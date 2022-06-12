//! A multi-value [`HeaderMap`] and its iterators.

use std::{borrow::Cow, collections::hash_map, ops};

use ahash::AHashMap;
use http::header::{HeaderName, HeaderValue};
use smallvec::{smallvec, SmallVec};

use crate::header::AsHeaderName;

/// A multi-map of HTTP headers.
///
/// `HeaderMap` is a "multi-map" of [`HeaderName`] to one or more [`HeaderValue`]s.
///
/// # Examples
/// ```
/// use actix_http::http::{header, HeaderMap, HeaderValue};
///
/// let mut map = HeaderMap::new();
///
/// map.insert(header::CONTENT_TYPE, HeaderValue::from_static("text/plain"));
/// map.insert(header::ORIGIN, HeaderValue::from_static("example.com"));
///
/// assert!(map.contains_key(header::CONTENT_TYPE));
/// assert!(map.contains_key(header::ORIGIN));
///
/// let mut removed = map.remove(header::ORIGIN);
/// assert_eq!(removed.next().unwrap(), "example.com");
///
/// assert!(!map.contains_key(header::ORIGIN));
/// ```
#[derive(Debug, Clone, Default)]
pub struct HeaderMap {
    pub(crate) inner: AHashMap<HeaderName, Value>,
}

/// A bespoke non-empty list for HeaderMap values.
#[derive(Debug, Clone)]
pub(crate) struct Value {
    inner: SmallVec<[HeaderValue; 4]>,
}

impl Value {
    fn one(val: HeaderValue) -> Self {
        Self {
            inner: smallvec![val],
        }
    }

    fn first(&self) -> &HeaderValue {
        &self.inner[0]
    }

    fn first_mut(&mut self) -> &mut HeaderValue {
        &mut self.inner[0]
    }

    fn append(&mut self, new_val: HeaderValue) {
        self.inner.push(new_val)
    }
}

impl ops::Deref for Value {
    type Target = SmallVec<[HeaderValue; 4]>;

    fn deref(&self) -> &Self::Target {
        &self.inner
    }
}

impl HeaderMap {
    /// Create an empty `HeaderMap`.
    ///
    /// The map will be created without any capacity; this function will not allocate.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::HeaderMap;
    /// let map = HeaderMap::new();
    ///
    /// assert!(map.is_empty());
    /// assert_eq!(0, map.capacity());
    /// ```
    pub fn new() -> Self {
        HeaderMap::default()
    }

    /// Create an empty `HeaderMap` with the specified capacity.
    ///
    /// The map will be able to hold at least `capacity` elements without needing to reallocate.
    /// If `capacity` is 0, the map will be created without allocating.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::HeaderMap;
    /// let map = HeaderMap::with_capacity(16);
    ///
    /// assert!(map.is_empty());
    /// assert!(map.capacity() >= 16);
    /// ```
    pub fn with_capacity(capacity: usize) -> Self {
        HeaderMap {
            inner: AHashMap::with_capacity(capacity),
        }
    }

    /// Create new `HeaderMap` from a `http::HeaderMap`-like drain.
    pub(crate) fn from_drain<I>(mut drain: I) -> Self
    where
        I: Iterator<Item = (Option<HeaderName>, HeaderValue)>,
    {
        let (first_name, first_value) = match drain.next() {
            None => return HeaderMap::new(),
            Some((name, val)) => {
                let name = name.expect("drained first item had no name");
                (name, val)
            }
        };

        let (lb, ub) = drain.size_hint();
        let capacity = ub.unwrap_or(lb);

        let mut map = HeaderMap::with_capacity(capacity);
        map.append(first_name.clone(), first_value);

        let (map, _) =
            drain.fold((map, first_name), |(mut map, prev_name), (name, value)| {
                let name = name.unwrap_or(prev_name);
                map.append(name.clone(), value);
                (map, name)
            });

        map
    }

    /// Returns the number of values stored in the map.
    ///
    /// See also: [`len_keys`](Self::len_keys).
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    /// assert_eq!(map.len(), 0);
    ///
    /// map.insert(header::ACCEPT, HeaderValue::from_static("text/plain"));
    /// map.insert(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// assert_eq!(map.len(), 2);
    ///
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    /// assert_eq!(map.len(), 3);
    /// ```
    pub fn len(&self) -> usize {
        self.inner
            .iter()
            .fold(0, |acc, (_, values)| acc + values.len())
    }

    /// Returns the number of _keys_ stored in the map.
    ///
    /// The number of values stored will be at least this number. See also: [`Self::len`].
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    /// assert_eq!(map.len_keys(), 0);
    ///
    /// map.insert(header::ACCEPT, HeaderValue::from_static("text/plain"));
    /// map.insert(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// assert_eq!(map.len_keys(), 2);
    ///
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    /// assert_eq!(map.len_keys(), 2);
    /// ```
    pub fn len_keys(&self) -> usize {
        self.inner.len()
    }

    /// Returns true if the map contains no elements.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    /// assert!(map.is_empty());
    ///
    /// map.insert(header::ACCEPT, HeaderValue::from_static("text/plain"));
    /// assert!(!map.is_empty());
    /// ```
    pub fn is_empty(&self) -> bool {
        self.inner.len() == 0
    }

    /// Clears the map, removing all name-value pairs.
    ///
    /// Keeps the allocated memory for reuse.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// map.insert(header::ACCEPT, HeaderValue::from_static("text/plain"));
    /// map.insert(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// assert_eq!(map.len(), 2);
    ///
    /// map.clear();
    /// assert!(map.is_empty());
    /// ```
    pub fn clear(&mut self) {
        self.inner.clear();
    }

    fn get_value(&self, key: impl AsHeaderName) -> Option<&Value> {
        match key.try_as_name(super::as_name::Seal).ok()? {
            Cow::Borrowed(name) => self.inner.get(name),
            Cow::Owned(name) => self.inner.get(&name),
        }
    }

    /// Returns a reference to the _first_ value associated with a header name.
    ///
    /// Returns `None` if there is no value associated with the key.
    ///
    /// Even when multiple values are associated with the key, the "first" one is returned but is
    /// not guaranteed to be chosen with any particular order; though, the returned item will be
    /// consistent for each call to `get` if the map has not changed.
    ///
    /// See also: [`get_all`](Self::get_all).
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// map.insert(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    ///
    /// let cookie = map.get(header::SET_COOKIE).unwrap();
    /// assert_eq!(cookie, "one=1");
    ///
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    /// assert_eq!(map.get(header::SET_COOKIE).unwrap(), "one=1");
    ///
    /// assert_eq!(map.get(header::SET_COOKIE), map.get("set-cookie"));
    /// assert_eq!(map.get(header::SET_COOKIE), map.get("Set-Cookie"));
    ///
    /// assert!(map.get(header::HOST).is_none());
    /// assert!(map.get("INVALID HEADER NAME").is_none());
    /// ```
    pub fn get(&self, key: impl AsHeaderName) -> Option<&HeaderValue> {
        self.get_value(key).map(Value::first)
    }

    /// Returns a mutable reference to the _first_ value associated a header name.
    ///
    /// Returns `None` if there is no value associated with the key.
    ///
    /// Even when multiple values are associated with the key, the "first" one is returned but is
    /// not guaranteed to be chosen with any particular order; though, the returned item will be
    /// consistent for each call to `get_mut` if the map has not changed.
    ///
    /// See also: [`get_all`](Self::get_all).
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// map.insert(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    ///
    /// let mut cookie = map.get_mut(header::SET_COOKIE).unwrap();
    /// assert_eq!(cookie, "one=1");
    ///
    /// *cookie = HeaderValue::from_static("three=3");
    /// assert_eq!(map.get(header::SET_COOKIE).unwrap(), "three=3");
    ///
    /// assert!(map.get(header::HOST).is_none());
    /// assert!(map.get("INVALID HEADER NAME").is_none());
    /// ```
    pub fn get_mut(&mut self, key: impl AsHeaderName) -> Option<&mut HeaderValue> {
        match key.try_as_name(super::as_name::Seal).ok()? {
            Cow::Borrowed(name) => self.inner.get_mut(name).map(Value::first_mut),
            Cow::Owned(name) => self.inner.get_mut(&name).map(Value::first_mut),
        }
    }

    /// Returns an iterator over all values associated with a header name.
    ///
    /// The returned iterator does not incur any allocations and will yield no items if there are no
    /// values associated with the key. Iteration order is **not** guaranteed to be the same as
    /// insertion order.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// let mut none_iter = map.get_all(header::ORIGIN);
    /// assert!(none_iter.next().is_none());
    ///
    /// map.insert(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    ///
    /// let mut set_cookies_iter = map.get_all(header::SET_COOKIE);
    /// assert_eq!(set_cookies_iter.next().unwrap(), "one=1");
    /// assert_eq!(set_cookies_iter.next().unwrap(), "two=2");
    /// assert!(set_cookies_iter.next().is_none());
    /// ```
    pub fn get_all(&self, key: impl AsHeaderName) -> GetAll<'_> {
        GetAll::new(self.get_value(key))
    }

    // TODO: get_all_mut ?

    /// Returns `true` if the map contains a value for the specified key.
    ///
    /// Invalid header names will simply return false.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    /// assert!(!map.contains_key(header::ACCEPT));
    ///
    /// map.insert(header::ACCEPT, HeaderValue::from_static("text/plain"));
    /// assert!(map.contains_key(header::ACCEPT));
    /// ```
    pub fn contains_key(&self, key: impl AsHeaderName) -> bool {
        match key.try_as_name(super::as_name::Seal) {
            Ok(Cow::Borrowed(name)) => self.inner.contains_key(name),
            Ok(Cow::Owned(name)) => self.inner.contains_key(&name),
            Err(_) => false,
        }
    }

    /// Inserts a name-value pair into the map.
    ///
    /// If the map already contained this key, the new value is associated with the key and all
    /// previous values are removed and returned as a `Removed` iterator. The key is not updated;
    /// this matters for types that can be `==` without being identical.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// map.insert(header::ACCEPT, HeaderValue::from_static("text/plain"));
    /// assert!(map.contains_key(header::ACCEPT));
    /// assert_eq!(map.len(), 1);
    ///
    /// let mut removed = map.insert(header::ACCEPT, HeaderValue::from_static("text/csv"));
    /// assert_eq!(removed.next().unwrap(), "text/plain");
    /// assert!(removed.next().is_none());
    ///
    /// assert_eq!(map.len(), 1);
    /// ```
    pub fn insert(&mut self, key: HeaderName, val: HeaderValue) -> Removed {
        let value = self.inner.insert(key, Value::one(val));
        Removed::new(value)
    }

    /// Inserts a name-value pair into the map.
    ///
    /// If the map already contained this key, the new value is added to the list of values
    /// currently associated with the key. The key is not updated; this matters for types that can
    /// be `==` without being identical.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// map.append(header::HOST, HeaderValue::from_static("example.com"));
    /// assert_eq!(map.len(), 1);
    ///
    /// map.append(header::ACCEPT, HeaderValue::from_static("text/csv"));
    /// assert_eq!(map.len(), 2);
    ///
    /// map.append(header::ACCEPT, HeaderValue::from_static("text/html"));
    /// assert_eq!(map.len(), 3);
    /// ```
    pub fn append(&mut self, key: HeaderName, value: HeaderValue) {
        match self.inner.entry(key) {
            hash_map::Entry::Occupied(mut entry) => {
                entry.get_mut().append(value);
            }
            hash_map::Entry::Vacant(entry) => {
                entry.insert(Value::one(value));
            }
        };
    }

    /// Removes all headers for a particular header name from the map.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("one=2"));
    ///
    /// assert_eq!(map.len(), 2);
    ///
    /// let mut removed = map.remove(header::SET_COOKIE);
    /// assert_eq!(removed.next().unwrap(), "one=1");
    /// assert_eq!(removed.next().unwrap(), "one=2");
    /// assert!(removed.next().is_none());
    ///
    /// assert!(map.is_empty());
    pub fn remove(&mut self, key: impl AsHeaderName) -> Removed {
        let value = match key.try_as_name(super::as_name::Seal) {
            Ok(Cow::Borrowed(name)) => self.inner.remove(name),
            Ok(Cow::Owned(name)) => self.inner.remove(&name),
            Err(_) => None,
        };

        Removed::new(value)
    }

    /// Returns the number of single-value headers the map can hold without needing to reallocate.
    ///
    /// Since this is a multi-value map, the actual capacity is much larger when considering
    /// each header name can be associated with an arbitrary number of values. The effect is that
    /// the size of `len` may be greater than `capacity` since it counts all the values.
    /// Conversely, [`len_keys`](Self::len_keys) will never be larger than capacity.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::HeaderMap;
    /// let map = HeaderMap::with_capacity(16);
    ///
    /// assert!(map.is_empty());
    /// assert!(map.capacity() >= 16);
    /// ```
    pub fn capacity(&self) -> usize {
        self.inner.capacity()
    }

    /// Reserves capacity for at least `additional` more headers to be inserted in the map.
    ///
    /// The header map may reserve more space to avoid frequent reallocations. Additional capacity
    /// only considers single-value headers.
    ///
    /// # Panics
    /// Panics if the new allocation size overflows usize.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::HeaderMap;
    /// let mut map = HeaderMap::with_capacity(2);
    /// assert!(map.capacity() >= 2);
    ///
    /// map.reserve(100);
    /// assert!(map.capacity() >= 102);
    ///
    /// assert!(map.is_empty());
    /// ```
    pub fn reserve(&mut self, additional: usize) {
        self.inner.reserve(additional)
    }

    /// An iterator over all name-value pairs.
    ///
    /// Names will be yielded for each associated value. So, if a key has 3 associated values, it
    /// will be yielded 3 times. The iteration order should be considered arbitrary.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// let mut iter = map.iter();
    /// assert!(iter.next().is_none());
    ///
    /// map.append(header::HOST, HeaderValue::from_static("duck.com"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    ///
    /// let mut iter = map.iter();
    /// assert!(iter.next().is_some());
    /// assert!(iter.next().is_some());
    /// assert!(iter.next().is_some());
    /// assert!(iter.next().is_none());
    ///
    /// let pairs = map.iter().collect::<Vec<_>>();
    /// assert!(pairs.contains(&(&header::HOST, &HeaderValue::from_static("duck.com"))));
    /// assert!(pairs.contains(&(&header::SET_COOKIE, &HeaderValue::from_static("one=1"))));
    /// assert!(pairs.contains(&(&header::SET_COOKIE, &HeaderValue::from_static("two=2"))));
    /// ```
    pub fn iter(&self) -> Iter<'_> {
        Iter::new(self.inner.iter())
    }

    /// An iterator over all contained header names.
    ///
    /// Each name will only be yielded once even if it has multiple associated values. The iteration
    /// order should be considered arbitrary.
    ///
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// let mut iter = map.keys();
    /// assert!(iter.next().is_none());
    ///
    /// map.append(header::HOST, HeaderValue::from_static("duck.com"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    ///
    /// let keys = map.keys().cloned().collect::<Vec<_>>();
    /// assert_eq!(keys.len(), 2);
    /// assert!(keys.contains(&header::HOST));
    /// assert!(keys.contains(&header::SET_COOKIE));
    /// ```
    pub fn keys(&self) -> Keys<'_> {
        Keys(self.inner.keys())
    }

    /// Clears the map, returning all name-value sets as an iterator.
    ///
    /// Header names will only be yielded for the first value in each set. All items that are
    /// yielded without a name and after an item with a name are associated with that same name.
    /// The first item will always contain a name.
    ///
    /// Keeps the allocated memory for reuse.
    /// # Examples
    /// ```
    /// # use actix_http::http::{header, HeaderMap, HeaderValue};
    /// let mut map = HeaderMap::new();
    ///
    /// let mut iter = map.drain();
    /// assert!(iter.next().is_none());
    /// drop(iter);
    ///
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("one=1"));
    /// map.append(header::SET_COOKIE, HeaderValue::from_static("two=2"));
    ///
    /// let mut iter = map.drain();
    /// assert_eq!(iter.next().unwrap(), (Some(header::SET_COOKIE), HeaderValue::from_static("one=1")));
    /// assert_eq!(iter.next().unwrap(), (None, HeaderValue::from_static("two=2")));
    /// drop(iter);
    ///
    /// assert!(map.is_empty());
    /// ```
    pub fn drain(&mut self) -> Drain<'_> {
        Drain::new(self.inner.drain())
    }
}

/// Note that this implementation will clone a [HeaderName] for each value.
impl IntoIterator for HeaderMap {
    type Item = (HeaderName, HeaderValue);
    type IntoIter = IntoIter;

    #[inline]
    fn into_iter(self) -> Self::IntoIter {
        IntoIter::new(self.inner.into_iter())
    }
}

impl<'a> IntoIterator for &'a HeaderMap {
    type Item = (&'a HeaderName, &'a HeaderValue);
    type IntoIter = Iter<'a>;

    #[inline]
    fn into_iter(self) -> Self::IntoIter {
        Iter::new(self.inner.iter())
    }
}

/// Iterator for all values with the same header name.
///
/// See [`HeaderMap::get_all`].
#[derive(Debug)]
pub struct GetAll<'a> {
    idx: usize,
    value: Option<&'a Value>,
}

impl<'a> GetAll<'a> {
    fn new(value: Option<&'a Value>) -> Self {
        Self { idx: 0, value }
    }
}

impl<'a> Iterator for GetAll<'a> {
    type Item = &'a HeaderValue;

    fn next(&mut self) -> Option<Self::Item> {
        let val = self.value?;

        match val.get(self.idx) {
            Some(val) => {
                self.idx += 1;
                Some(val)
            }
            None => {
                // current index is none; remove value to fast-path future next calls
                self.value = None;
                None
            }
        }
    }

    fn size_hint(&self) -> (usize, Option<usize>) {
        match self.value {
            Some(val) => (val.len(), Some(val.len())),
            None => (0, Some(0)),
        }
    }
}

/// Iterator for owned [`HeaderValue`]s with the same associated [`HeaderName`] returned from methods
/// on [`HeaderMap`] that remove or replace items.
#[derive(Debug)]
pub struct Removed {
    inner: Option<smallvec::IntoIter<[HeaderValue; 4]>>,
}

impl<'a> Removed {
    fn new(value: Option<Value>) -> Self {
        let inner = value.map(|value| value.inner.into_iter());
        Self { inner }
    }
}

impl Iterator for Removed {
    type Item = HeaderValue;

    #[inline]
    fn next(&mut self) -> Option<Self::Item> {
        self.inner.as_mut()?.next()
    }

    #[inline]
    fn size_hint(&self) -> (usize, Option<usize>) {
        match self.inner {
            Some(ref iter) => iter.size_hint(),
            None => (0, None),
        }
    }
}

/// Iterator over all [`HeaderName`]s in the map.
#[derive(Debug)]
pub struct Keys<'a>(hash_map::Keys<'a, HeaderName, Value>);

impl<'a> Iterator for Keys<'a> {
    type Item = &'a HeaderName;

    #[inline]
    fn next(&mut self) -> Option<Self::Item> {
        self.0.next()
    }

    #[inline]
    fn size_hint(&self) -> (usize, Option<usize>) {
        self.0.size_hint()
    }
}

#[derive(Debug)]
pub struct Iter<'a> {
    inner: hash_map::Iter<'a, HeaderName, Value>,
    multi_inner: Option<(&'a HeaderName, &'a SmallVec<[HeaderValue; 4]>)>,
    multi_idx: usize,
}

impl<'a> Iter<'a> {
    fn new(iter: hash_map::Iter<'a, HeaderName, Value>) -> Self {
        Self {
            inner: iter,
            multi_idx: 0,
            multi_inner: None,
        }
    }
}

impl<'a> Iterator for Iter<'a> {
    type Item = (&'a HeaderName, &'a HeaderValue);

    fn next(&mut self) -> Option<Self::Item> {
        // handle in-progress multi value lists first
        if let Some((ref name, ref mut vals)) = self.multi_inner {
            match vals.get(self.multi_idx) {
                Some(val) => {
                    self.multi_idx += 1;
                    return Some((name, val));
                }
                None => {
                    // no more items in value list; reset state
                    self.multi_idx = 0;
                    self.multi_inner = None;
                }
            }
        }

        let (name, value) = self.inner.next()?;

        // set up new inner iter and recurse into it
        self.multi_inner = Some((name, &value.inner));
        self.next()
    }

    #[inline]
    fn size_hint(&self) -> (usize, Option<usize>) {
        // take inner lower bound
        // make no attempt at an upper bound
        (self.inner.size_hint().0, None)
    }
}

/// Iterator over drained name-value pairs.
///
/// Iterator items are `(Option<HeaderName>, HeaderValue)` to avoid cloning.
#[derive(Debug)]
pub struct Drain<'a> {
    inner: hash_map::Drain<'a, HeaderName, Value>,
    multi_inner: Option<(Option<HeaderName>, SmallVec<[HeaderValue; 4]>)>,
    multi_idx: usize,
}

impl<'a> Drain<'a> {
    fn new(iter: hash_map::Drain<'a, HeaderName, Value>) -> Self {
        Self {
            inner: iter,
            multi_inner: None,
            multi_idx: 0,
        }
    }
}

impl<'a> Iterator for Drain<'a> {
    type Item = (Option<HeaderName>, HeaderValue);

    fn next(&mut self) -> Option<Self::Item> {
        // handle in-progress multi value iterators first
        if let Some((ref mut name, ref mut vals)) = self.multi_inner {
            if !vals.is_empty() {
                // OPTIMIZE: array removals
                return Some((name.take(), vals.remove(0)));
            } else {
                // no more items in value iterator; reset state
                self.multi_inner = None;
                self.multi_idx = 0;
            }
        }

        let (name, value) = self.inner.next()?;

        // set up new inner iter and recurse into it
        self.multi_inner = Some((Some(name), value.inner));
        self.next()
    }

    #[inline]
    fn size_hint(&self) -> (usize, Option<usize>) {
        // take inner lower bound
        // make no attempt at an upper bound
        (self.inner.size_hint().0, None)
    }
}

/// Iterator over owned name-value pairs.
///
/// Implementation necessarily clones header names for each value.
#[derive(Debug)]
pub struct IntoIter {
    inner: hash_map::IntoIter<HeaderName, Value>,
    multi_inner: Option<(HeaderName, smallvec::IntoIter<[HeaderValue; 4]>)>,
}

impl IntoIter {
    fn new(inner: hash_map::IntoIter<HeaderName, Value>) -> Self {
        Self {
            inner,
            multi_inner: None,
        }
    }
}

impl Iterator for IntoIter {
    type Item = (HeaderName, HeaderValue);

    fn next(&mut self) -> Option<Self::Item> {
        // handle in-progress multi value iterators first
        if let Some((ref name, ref mut vals)) = self.multi_inner {
            match vals.next() {
                Some(val) => {
                    return Some((name.clone(), val));
                }
                None => {
                    // no more items in value iterator; reset state
                    self.multi_inner = None;
                }
            }
        }

        let (name, value) = self.inner.next()?;

        // set up new inner iter and recurse into it
        self.multi_inner = Some((name, value.inner.into_iter()));
        self.next()
    }

    #[inline]
    fn size_hint(&self) -> (usize, Option<usize>) {
        // take inner lower bound
        // make no attempt at an upper bound
        (self.inner.size_hint().0, None)
    }
}

#[cfg(test)]
mod tests {
    use http::header;

    use super::*;

    #[test]
    fn create() {
        let map = HeaderMap::new();
        assert_eq!(map.len(), 0);
        assert_eq!(map.capacity(), 0);

        let map = HeaderMap::with_capacity(16);
        assert_eq!(map.len(), 0);
        assert!(map.capacity() >= 16);
    }

    #[test]
    fn insert() {
        let mut map = HeaderMap::new();

        map.insert(header::LOCATION, HeaderValue::from_static("/test"));
        assert_eq!(map.len(), 1);
    }

    #[test]
    fn contains() {
        let mut map = HeaderMap::new();
        assert!(!map.contains_key(header::LOCATION));

        map.insert(header::LOCATION, HeaderValue::from_static("/test"));
        assert!(map.contains_key(header::LOCATION));
        assert!(map.contains_key("Location"));
        assert!(map.contains_key("Location".to_owned()));
        assert!(map.contains_key("location"));
    }

    #[test]
    fn entries_iter() {
        let mut map = HeaderMap::new();

        map.append(header::HOST, HeaderValue::from_static("duck.com"));
        map.append(header::COOKIE, HeaderValue::from_static("one=1"));
        map.append(header::COOKIE, HeaderValue::from_static("two=2"));

        let mut iter = map.iter();
        assert!(iter.next().is_some());
        assert!(iter.next().is_some());
        assert!(iter.next().is_some());
        assert!(iter.next().is_none());

        let pairs = map.iter().collect::<Vec<_>>();
        assert!(pairs.contains(&(&header::HOST, &HeaderValue::from_static("duck.com"))));
        assert!(pairs.contains(&(&header::COOKIE, &HeaderValue::from_static("one=1"))));
        assert!(pairs.contains(&(&header::COOKIE, &HeaderValue::from_static("two=2"))));
    }

    #[test]
    fn drain_iter() {
        let mut map = HeaderMap::new();

        map.append(header::COOKIE, HeaderValue::from_static("one=1"));
        map.append(header::COOKIE, HeaderValue::from_static("two=2"));

        let mut vals = vec![];
        let mut iter = map.drain();

        let (name, val) = iter.next().unwrap();
        assert_eq!(name, Some(header::COOKIE));
        vals.push(val);

        let (name, val) = iter.next().unwrap();
        assert!(name.is_none());
        vals.push(val);

        assert!(vals.contains(&HeaderValue::from_static("one=1")));
        assert!(vals.contains(&HeaderValue::from_static("two=2")));

        assert!(iter.next().is_none());
        drop(iter);

        assert!(map.is_empty());
    }

    #[test]
    fn entries_into_iter() {
        let mut map = HeaderMap::new();

        map.append(header::HOST, HeaderValue::from_static("duck.com"));
        map.append(header::COOKIE, HeaderValue::from_static("one=1"));
        map.append(header::COOKIE, HeaderValue::from_static("two=2"));

        let mut iter = map.into_iter();
        assert!(iter.next().is_some());
        assert!(iter.next().is_some());
        assert!(iter.next().is_some());
        assert!(iter.next().is_none());
    }

    #[test]
    fn iter_and_into_iter_same_order() {
        let mut map = HeaderMap::new();

        map.append(header::HOST, HeaderValue::from_static("duck.com"));
        map.append(header::COOKIE, HeaderValue::from_static("one=1"));
        map.append(header::COOKIE, HeaderValue::from_static("two=2"));

        let mut iter = map.iter();
        let mut into_iter = map.clone().into_iter();

        assert_eq!(iter.next().map(owned_pair), into_iter.next());
        assert_eq!(iter.next().map(owned_pair), into_iter.next());
        assert_eq!(iter.next().map(owned_pair), into_iter.next());
        assert_eq!(iter.next().map(owned_pair), into_iter.next());
    }

    #[test]
    fn get_all_and_remove_same_order() {
        let mut map = HeaderMap::new();

        map.append(header::COOKIE, HeaderValue::from_static("one=1"));
        map.append(header::COOKIE, HeaderValue::from_static("two=2"));

        let mut vals = map.get_all(header::COOKIE);
        let mut removed = map.clone().remove(header::COOKIE);

        assert_eq!(vals.next(), removed.next().as_ref());
        assert_eq!(vals.next(), removed.next().as_ref());
        assert_eq!(vals.next(), removed.next().as_ref());
    }

    fn owned_pair<'a>(
        (name, val): (&'a HeaderName, &'a HeaderValue),
    ) -> (HeaderName, HeaderValue) {
        (name.clone(), val.clone())
    }
}