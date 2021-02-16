# https://docs.bazel.build/versions/master/skylark/lib/range.html

nums = []
nums.append(1)

def print_seq():
    ns = []
    # for loops are not allowed at the top level
    # must reside in a function or list comprehension
    for x in range(1, 10, 3):
        ns.append(x)
        # trying to mutate a frozen list value
        # https://docs.bazel.build/versions/master/skylark/language.html#mutability
        # nums.append(x)
    print("list =", ns)

# runs on load
print("str:", "starlark rocks")
print("tuple:", (1,2,3))
print("map:", {
    "foo": 1,
    "bar": 2,
})
print("int:", 2^3)

# https://docs.bazel.build/versions/master/skylark/language.html
def fizz_buzz(n):
  """Print Fizz Buzz numbers from 1 to n."""
  for i in range(1, n + 1):
    s = ""
    if i % 3 == 0:
      s += "Fizz"
    if i % 5 == 0:
      s += "Buzz"
    print(s if s else i)
