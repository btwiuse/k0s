// #[macro_use]
// extern crate serde_derive;
extern crate tokio;
extern crate schemars;
extern crate kube;
// extern crate kube_derive;
extern crate kube_runtime;
extern crate futures;

use kube::Client;
use kube::api::Api;
use kube::api::Meta;
use kube::api::ListParams;
use kube_runtime::watcher;
use kube_runtime::watcher::Event;
use futures::StreamExt;
use futures::TryStreamExt;

// macros
use kube::CustomResource;
use serde::Serialize;
use serde::Deserialize;
use schemars::JsonSchema;

// this is our new Book struct
#[derive(CustomResource, Serialize, Deserialize, Default, Clone, Debug, PartialEq, JsonSchema)]
#[kube(group = "example.technosophos.com", version = "v1", kind = "Book", namespaced)]
pub struct KubeBook {
    pub title: String,
    pub authors: Option<Vec<String>>,
}

#[tokio::main]
async fn main() -> (){
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
            Event::Applied(b) => println!("Created({}): Title: {}", Meta::name(&b), b.spec.title),
            Event::Deleted(b) => println!("Deleted({}): Title: {}", Meta::name(&b), b.spec.title),
            _ => (), // Ignore Restarted (We already use list)
        };
    }
}
