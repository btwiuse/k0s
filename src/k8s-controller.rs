// #[macro_use]
// extern crate serde_derive;
extern crate kube;
extern crate schemars;
extern crate tokio;
// extern crate kube_derive;
extern crate futures;
extern crate kube_runtime;

use futures::StreamExt;
use futures::TryStreamExt;
use kube::api::Api;
use kube::api::ListParams;
// use kube::api::Meta;
use kube::Client;
use kube_runtime::watcher;
use kube_runtime::watcher::Event;

// macros
use kube::CustomResource;
use schemars::JsonSchema;
use serde::Deserialize;
use serde::Serialize;

// this is our new Book struct
#[derive(CustomResource, Serialize, Deserialize, Default, Clone, Debug, PartialEq, JsonSchema)]
#[kube(
    group = "example.technosophos.com",
    version = "v1",
    kind = "Book",
    namespaced
)]
pub struct KubeBook {
    pub title: String,
    pub authors: Option<Vec<String>>,
}

#[tokio::main]
async fn main() -> () {
    // create a new client
    let client = Client::try_default().await.unwrap();

    // Describe the CRD we're working with
    // This is basically the fields from our CRD definition.
    let books = Api::<Book>::namespaced(client, "default");

    let lp = ListParams::default();

    // create our informer and start listening
    let mut w = watcher::watcher(books, lp).boxed();
    while let Some(event) = w.try_next().await.unwrap() {
        match event {
            Event::Applied(b) => println!("Created({}): Title: {}", b.metadata.name.unwrap(), b.spec.title),
            Event::Deleted(b) => println!("Deleted({}): Title: {}", b.metadata.name.unwrap(), b.spec.title),
            _ => (), // Ignore Restarted (We already use list)
        };
    }
}
