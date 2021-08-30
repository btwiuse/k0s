use actix::prelude::*;
use actix::{Actor, ActorContext, Addr, AsyncContext, Handler, StreamHandler};
use actix_web::{
    get, middleware, web, App, Error, HttpRequest, HttpResponse, HttpServer, Responder,
};
use actix_web_actors::ws;
use lazy_static::lazy_static;
use serde::Deserialize;
use simple_logger::SimpleLogger;
use std::collections::{HashMap};
use std::sync::{Arc, RwLock};

lazy_static! {
    static ref AGENTS: Arc<RwLock<HashMap<String, Addr<Agent>>>> = Arc::new(RwLock::new(HashMap::<String, Addr<Agent>>::new()));
    static ref PENDING_CLIENTS: Arc<RwLock<HashMap<String, Addr<Client>>>> = Arc::new(RwLock::new(HashMap::<String, Addr<Client>>::new()));
    static ref CLIENT_TERMINAL_PAIRS: Arc<RwLock<HashMap<Addr<Client>, Addr<Terminal>>>> = Arc::new(RwLock::new(HashMap::<Addr<Client>, Addr<Terminal>>::new()));
}

struct Agent {
    // seq: u32,
    // agent_id: String,
    agent_id: Option<String>,
}

#[derive(Message)]
#[rtype(result = "()")]
struct NewTerminal;

impl Handler<NewTerminal> for Agent {
    type Result = ();
    fn handle(&mut self, _msg: NewTerminal, ctx: &mut Self::Context) {
        // self.handle_msg(msg.0, ctx);
        println!("received NewTerminal");
        ctx.binary("TERMINAL\n");
    }
}

impl Agent {
    fn is_initialized(&mut self) -> bool {
        // self.seq == 0
        self.agent_id.is_some()
    }
    fn on_init(&mut self, ctx: &mut <Self as Actor>::Context) {
        let mut agents = AGENTS.write().unwrap();
        let id = self.agent_id.as_ref().unwrap().clone();
        // let set = agents.entry(id).or_insert(HashSet::<Addr<Self>>::new());
        // set.insert(ctx.address());
        // agents[id] = ctx.address();
        *agents.entry(id).or_insert(ctx.address()) = ctx.address();
    }
    fn on_stop(&mut self, _ctx: &mut <Self as Actor>::Context) {
        let mut agents = AGENTS.write().unwrap();
        let id = self.agent_id.as_ref().unwrap().clone();
        // let set = agents.entry(id).or_insert(HashSet::<Addr<Self>>::new());
        // set.remove(&ctx.address());
        agents.remove(&id);
    }
    fn init(&mut self, msg: ws::Message, ctx: &mut <Self as Actor>::Context) {
        let mut agent_info = "".to_string();
        match msg {
            ws::Message::Text(m) => {
                println!("agent_info text: {:?}", m);
                agent_info = m.to_string();
                // ctx.text(m);
            }
            ws::Message::Binary(m) => {
                println!("agent_info binary: {:?}", m);
                agent_info = String::from_utf8_lossy(&m).to_string();
                // ctx.binary(m);
            }
            _ => (),
        };
        println!("agent_info: {}", agent_info);
        // let deserialized: Result<ID, _> = serde_json::from_str::<ID>(&agent_info);
        let deserialized = serde_json::from_str::<ID>(&agent_info);
        if let Ok(id) = deserialized {
            if let Some(id) = id.id {
                println!("id: {}", id);
                self.agent_id = Some(id);
                self.on_init(ctx);
            }
        }
    }
    fn handle_cmd(&mut self, msg: ws::Message, _ctx: &mut <Self as Actor>::Context) {
        let mut cmd = "".to_string();
        match msg {
            ws::Message::Text(m) => {
                println!("agent_info text: {:?}", m);
                cmd = m.to_string();
                // ctx.text(m);
            }
            ws::Message::Binary(m) => {
                println!("agent_info binary: {:?}", m);
                cmd = String::from_utf8_lossy(&m).to_string();
                // ctx.binary(m);
            }
            _ => (),
        };
        println!("cmd: {}", cmd);
    }
}

impl Actor for Agent {
    type Context = ws::WebsocketContext<Self>;
    fn stopped(&mut self, ctx: &mut Self::Context) {
        self.on_stop(ctx)
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for Agent {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        if self.is_initialized() {
            // terminal output
            self.handle_cmd(msg.unwrap(), ctx)
        } else {
            // msg contains json line formatted agent info, hopefully
            self.init(msg.unwrap(), ctx);
            // ctx.binary("TERMINAL\n");
        }
    }
}

#[get("/api/rpc")]
async fn api_agent(
    req: HttpRequest,
    stream: web::Payload,
    id: web::Query<ID>,
) -> Result<HttpResponse, Error> {
    // let id = id.id.clone();
    // let id = "todo"
    // let seq = agent_count(id);
    let handler = Agent {
        // seq: 0,
        // agent_id: id,
        agent_id: id.id.clone(),
    };

    ws::start(handler, &req, stream)
}

struct Terminal {
    agent_id: String,
    client_addr: Option<Addr<Client>>,
}

#[derive(Message)]
#[rtype(result = "()")]
struct WsMessage(ws::Message);

impl Handler<WsMessage> for Terminal {
    type Result = ();
    fn handle(&mut self, msg: WsMessage, ctx: &mut Self::Context) {
        self.handle_msg(msg.0, ctx);
    }
}

impl Actor for Terminal {
    type Context = ws::WebsocketContext<Self>;
    fn started(&mut self, ctx: &mut Self::Context) {
        let mut pending_clients = PENDING_CLIENTS.write().unwrap();
        let id = self.agent_id.clone();
        let client_addr = pending_clients[&id].clone();
        self.client_addr = Some(client_addr.clone());
        pending_clients.remove(&id);

        let mut client_terminal_pairs = CLIENT_TERMINAL_PAIRS.write().unwrap();
        *client_terminal_pairs
            .entry(client_addr)
            .or_insert(ctx.address()) = ctx.address();

        println!("agent started");
    }
    fn stopped(&mut self, _ctx: &mut Self::Context) {
        // close client too
    }
}

impl Terminal {
    fn forward(&mut self, msg: ws::Message, _ctx: &mut <Self as Actor>::Context) {
        // println!("agent forward");
        if let Some(client_addr) = self.client_addr.clone() {
            client_addr.do_send(WsMessage(msg))
        }
    }
    fn handle_msg(&mut self, msg: ws::Message, ctx: &mut <Self as Actor>::Context) {
        let mut message = "".to_string();
        match msg {
            ws::Message::Text(m) => {
                message = m.to_string();
                ctx.text(m);
            }
            ws::Message::Binary(m) => {
                message = String::from_utf8_lossy(&m).to_string();
                ctx.binary(m);
            }
            _ => (),
        };
        println!("terminal::handle_msg({}): {}", message.len(), message);
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for Terminal {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(m) => {
                self.forward(m, ctx);
            }
            _ => {}
        }
    }
}

#[get("/api/terminal")]
async fn api_terminal(
    req: HttpRequest,
    stream: web::Payload,
    id: web::Query<ID>,
) -> Result<HttpResponse, Error> {
    // let id = id.id.clone();
    // let id = "todo"
    // let seq = agent_count(id);
    let agent_id = id.id.clone().unwrap();
    let handler = Terminal {
        agent_id: agent_id.clone(),
        client_addr: None,
    };

    println!("/api/terminal?id={}", agent_id);

    ws::start(handler, &req, stream)
}

struct Client {
    agent_id: String,
    agent_addr: Addr<Agent>,
}

impl Handler<WsMessage> for Client {
    type Result = ();
    fn handle(&mut self, msg: WsMessage, ctx: &mut Self::Context) {
        self.handle_msg(msg.0, ctx);
    }
}

impl Actor for Client {
    type Context = ws::WebsocketContext<Self>;
    fn started(&mut self, ctx: &mut Self::Context) {
        let mut pending_clients = PENDING_CLIENTS.write().unwrap();
        let id = self.agent_id.clone();
        *pending_clients.entry(id).or_insert(ctx.address()) = ctx.address();

        self.agent_addr.do_send(NewTerminal);
        println!("client started");
    }
    fn stopped(&mut self, _ctx: &mut Self::Context) {}
}

impl Client {
    fn forward(&mut self, msg: ws::Message, ctx: &mut <Self as Actor>::Context) {
        // println!("client forward");
        loop {
            let client_terminal_pairs = CLIENT_TERMINAL_PAIRS.read().unwrap();
            if client_terminal_pairs.get(&ctx.address()).is_some() {
                break;
            } else {
                std::thread::sleep(std::time::Duration::from_millis(10));
            }
        }
        let client_terminal_pairs = CLIENT_TERMINAL_PAIRS.read().unwrap();
        let terminal_addr = client_terminal_pairs.get(&ctx.address()).unwrap();
        terminal_addr.do_send(WsMessage(msg));
    }
    fn handle_msg(&mut self, msg: ws::Message, ctx: &mut <Self as Actor>::Context) {
        let mut message = "".to_string();
        match msg {
            ws::Message::Text(m) => {
                message = m.to_string();
                ctx.text(m);
            }
            ws::Message::Binary(m) => {
                message = String::from_utf8_lossy(&m).to_string();
                println!("bin({}): {:?}", m.len(), m);
                ctx.binary(m);
            }
            _ => (),
        };
        println!("client::handle_msg({}): {}", message.len(), message);
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for Client {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(m) => {
                self.forward(m, ctx);
            }
            _ => {}
        }
    }
}

#[get("/api/agent/{id}/terminal")]
async fn api_client(req: HttpRequest, stream: web::Payload, id: web::Path<ID>) -> impl Responder {
    println!("path: {}\n", req.path());

    let agent_id = id.id.clone().unwrap();
    let agents = AGENTS.read().unwrap();
    let agent_addr = agents.get(&agent_id);
    if agent_addr.is_none() {
        return Ok(HttpResponse::NotFound().body("not found"));
    }
    let handler = Client {
        agent_id: agent_id,
        agent_addr: agent_addr.unwrap().clone(),
    };
    ws::start(handler, &req, stream)
}

#[derive(Deserialize)]
struct ID {
    id: Option<String>,
}

struct Sink;

impl Actor for Sink {
    type Context = ws::WebsocketContext<Self>;
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for Sink {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(ws::Message::Ping(msg)) => ctx.pong(&msg),
            Ok(ws::Message::Text(text)) => {
                if text == "q" {
                    ctx.stop()
                } else {
                    println!("text: {}", text)
                }
            }
            Ok(ws::Message::Binary(bin)) => {
                let data = String::from_utf8_lossy(&bin);
                println!("{}", data)
            }
            _ => (),
        }
    }
}

#[derive(Clone)]
struct AppState {
    counter: usize,
}

impl AppState {
    fn _inc(&mut self) -> usize {
        self.counter += 1;
        self.counter
    }
}

/// Telemetry entry point. Listening by default on 127.0.0.1:8000.
/// This can be changed using the `PORT` and `BIND` ENV variables.
#[actix_web::main]
async fn main() -> std::io::Result<()> {
    SimpleLogger::new()
        .with_level(log::LevelFilter::Info)
        .init()
        .expect("Must be able to start a logger");
    log::info!("Starting telemetry version: {}", env!("CARGO_PKG_VERSION"));
    let app_state = web::Data::new(RwLock::new(AppState { counter: 0 }));
    HttpServer::new(move || {
        App::new()
            .wrap(middleware::NormalizePath::default())
            .app_data(app_state.clone())
            .service(api_client)
            .service(api_terminal)
            .service(api_agent)
    })
    .bind("127.0.0.1:8000")?
    .run()
    .await
}
