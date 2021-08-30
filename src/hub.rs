use actix::prelude::*;
use actix_web::*;
use actix_web_actors::ws;
use lazy_static::lazy_static;
use serde::Deserialize;
use simple_logger::SimpleLogger;
use std::collections::HashMap;
use std::sync::{Arc, RwLock};

lazy_static! {
    // static ref X_AGENTS: Arc<RwLock<HashMap<String, Addr<Agent>>>> = Arc::new(RwLock::new(HashMap::<String, Addr<Agent>>::new()));
    static ref X_PENDING_CLIENTS: Arc<RwLock<HashMap<String, Addr<Client>>>> =
        Arc::new(RwLock::new(HashMap::<String, Addr<Client>>::new()));
    static ref X_CLIENT_TERMINAL_PAIRS: Arc<RwLock<HashMap<Addr<Client>, Addr<Terminal>>>> =
        Arc::new(RwLock::new(HashMap::<Addr<Client>, Addr<Terminal>>::new()));
}

struct Agent {
    // seq: u32,
    // agent_id: String,
    agent_id: Option<String>,
    app_store: Addr<AppStore>,
}

impl Agent {
    fn new(agent_id: Option<String>, app_store: Addr<AppStore>) -> Self {
        Agent {
            agent_id,
            app_store,
        }
    }
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
        let id = self.agent_id.as_ref().unwrap().clone();
        /*
        let mut agents = X_AGENTS.write().unwrap();
        // let set = agents.entry(id).or_insert(HashSet::<Addr<Self>>::new());
        // set.insert(ctx.address());
        // agents[id] = ctx.address();
        *agents.entry(id).or_insert(ctx.address()) = ctx.address();
        */
        self.app_store.do_send(Event::AddAgent(id, ctx.address()));
    }
    fn on_stop(&mut self, _ctx: &mut <Self as Actor>::Context) {
        let id = self.agent_id.as_ref().unwrap().clone();
        /*
        let mut agents = X_AGENTS.write().unwrap();
        // let set = agents.entry(id).or_insert(HashSet::<Addr<Self>>::new());
        // set.remove(&ctx.address());
        agents.remove(&id);
        */
        self.app_store.do_send(Event::RemoveAgent(id));
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
                println!("env ID={} ~/kubotnetes/k0s.io miniclient :8000", id);
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
    agent_id: web::Query<ID>,
    // data: web::Data<RwLock<AppState>>,
    store: web::Data<Addr<AppStore>>,
) -> Result<HttpResponse, Error> {
    // let id = id.id.clone();
    // let id = "todo"
    // let seq = agent_count(id);
    // let mut data = data.write().unwrap();
    // data.counter += 1;

    let agent_id = agent_id.id.clone();
    let store = store.get_ref();
    let handler = Agent::new(agent_id, store.clone());

    store.do_send(Event::Inc);

    ws::start(handler, &req, stream)
}

struct Terminal {
    agent_id: String,
    client_addr: Option<Addr<Client>>,
    app_store: Addr<AppStore>,
}

impl Terminal {
    fn new(agent_id: String, client_addr: Option<Addr<Client>>, app_store: Addr<AppStore>) -> Self {
        Terminal {
            agent_id,
            client_addr,
            app_store,
        }
    }
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
        let mut pending_clients = X_PENDING_CLIENTS.write().unwrap();
        let id = self.agent_id.clone();
        let client_addr = pending_clients[&id].clone();
        self.client_addr = Some(client_addr.clone());
        pending_clients.remove(&id);

        let mut client_terminal_pairs = X_CLIENT_TERMINAL_PAIRS.write().unwrap();
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
    agent_id: web::Query<ID>,
    store: web::Data<Addr<AppStore>>,
) -> Result<HttpResponse, Error> {
    // let id = id.id.clone();
    // let id = "todo"
    // let seq = agent_count(id);
    let agent_id = agent_id.id.clone().unwrap();
    let client_addr: Option<Addr<Client>> = None;
    let store = store.get_ref();
    let handler = Terminal::new(agent_id.clone(), client_addr, store.clone());

    println!("/api/terminal?id={}", agent_id);

    ws::start(handler, &req, stream)
}

struct Client {
    agent_id: String,
    agent_addr: Addr<Agent>,
    app_store: Addr<AppStore>,
}

impl Client {
    fn new(agent_id: String, agent_addr: Addr<Agent>, app_store: Addr<AppStore>) -> Self {
        Client {
            agent_id,
            agent_addr,
            app_store,
        }
    }
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
        let mut pending_clients = X_PENDING_CLIENTS.write().unwrap();
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
            let client_terminal_pairs = X_CLIENT_TERMINAL_PAIRS.read().unwrap();
            if client_terminal_pairs.get(&ctx.address()).is_some() {
                break;
            } else {
                std::thread::sleep(std::time::Duration::from_millis(10));
            }
        }
        let client_terminal_pairs = X_CLIENT_TERMINAL_PAIRS.read().unwrap();
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
async fn api_client(
    req: HttpRequest,
    stream: web::Payload,
    id: web::Path<ID>,
    store: web::Data<Addr<AppStore>>,
) -> impl Responder {
    println!("path: {}\n", req.path());

    let agent_id = id.id.clone().unwrap();
    /*
    let agents = X_AGENTS.read().unwrap();
    let agent_addr = agents.get(&agent_id);
    let contains = agent_addr.is_none()
    */
    let (contains, agent_addr) = store.send(ContainsAgent(agent_id.clone())).await.unwrap();
    if !contains {
        return Ok(HttpResponse::NotFound().body("not found"));
    }
    let store = store.get_ref();
    let handler = Client::new(agent_id, agent_addr.unwrap().clone(), store.clone());
    ws::start(handler, &req, stream)
}

#[derive(Deserialize)]
struct ID {
    id: Option<String>,
}

#[derive(Clone)]
struct AppStore {
    counter: usize,
    agents: HashMap<String, Addr<Agent>>,
    pending_clients: HashMap<String, Addr<Client>>,
    client_terminal_pairs: HashMap<Addr<Client>, Addr<Terminal>>,
}

impl Actor for AppStore {
    type Context = Context<Self>;
}

impl AppStore {
    fn new() -> Self {
        AppStore {
            counter: 0,
            agents: HashMap::<String, Addr<Agent>>::new(),
            pending_clients: HashMap::<String, Addr<Client>>::new(),
            client_terminal_pairs: HashMap::<Addr<Client>, Addr<Terminal>>::new(),
        }
    }
}

#[derive(Message)]
#[rtype(result = "()")]
enum Event {
    Inc,
    AddAgent(String, Addr<Agent>),
    RemoveAgent(String),
}

impl Handler<Event> for AppStore {
    type Result = ();
    fn handle(&mut self, e: Event, _ctx: &mut Self::Context) {
        match e {
            Event::Inc => self.inc(),
            Event::AddAgent(id, addr) => self.add_agent(id, addr),
            Event::RemoveAgent(id) => self.remove_agent(id),
        }
    }
}

#[derive(Message)]
#[rtype(result = "(bool, Option<Addr<Agent>>)")]
struct ContainsAgent(String);

impl Handler<ContainsAgent> for AppStore {
    // type Result = (bool, Option<Addr<Agent>>);
    type Result = MessageResult<ContainsAgent>;
    fn handle(&mut self, c: ContainsAgent, _ctx: &mut Self::Context) -> Self::Result {
        MessageResult(self.has_agent(c.0))
    }
}

impl AppStore {
    fn inc(&mut self) {
        self.counter += 1;
        println!("store counter: {}", self.counter);
    }
    fn add_agent(&mut self, agent_id: String, agent_addr: Addr<Agent>) {
        *self
            .agents
            .entry(agent_id.clone())
            .or_insert(agent_addr.clone()) = agent_addr.clone();
        println!("store agents added: {}", agent_id);
    }
    fn remove_agent(&mut self, agent_id: String) {
        self.agents.remove(&agent_id);
        println!("store agents removed: {}", agent_id);
    }
    fn has_agent(&mut self, agent_id: String) -> (bool, Option<Addr<Agent>>) {
        let agent_addr = self.agents.get(&agent_id);
        let contains = !agent_addr.is_none();
        println!("store agents contains {}: {}", agent_id, contains);
        (contains, agent_addr.cloned())
    }
}

#[get("/inc")]
async fn inc(store: web::Data<Addr<AppStore>>) -> impl Responder {
    store.do_send(Event::Inc);
    format!("{}", "Ok")
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
    let store = AppStore::new().start();
    HttpServer::new(move || {
        #[allow(deprecated)]
        App::new()
            .wrap(middleware::NormalizePath::default())
            // compiler said should use app_data here. but didn't work for me
            .data(store.clone())
            .service(inc)
            .service(api_client)
            .service(api_terminal)
            .service(api_agent)
    })
    .bind("127.0.0.1:8000")?
    .run()
    .await
}
