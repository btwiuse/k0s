use actix::prelude::{
    Actor, ActorContext, Addr, AsyncContext, Context, Handler, Message, MessageResult,
    StreamHandler,
};
use actix_http::ws::{Codec, Item};
use actix_web::{
    get, middleware, web, App, Error, HttpRequest, HttpResponse, HttpServer, Responder,
};
use actix_web_actors::ws;
use bytes::{BufMut, BytesMut};
use serde::Deserialize;
use simple_logger::SimpleLogger;
use std::collections::{HashMap, VecDeque};

#[derive(Debug)]
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
struct Stop;

impl Handler<Stop> for Terminal {
    type Result = ();
    fn handle(&mut self, _s: Stop, ctx: &mut Self::Context) {
        ctx.stop();
    }
}

impl Handler<Stop> for Client {
    type Result = ();
    fn handle(&mut self, _s: Stop, ctx: &mut Self::Context) {
        ctx.stop();
    }
}

#[derive(Message)]
#[rtype(result = "()")]
struct NewTunnel(Command);

impl Handler<NewTunnel> for Agent {
    type Result = ();
    fn handle(&mut self, n: NewTunnel, ctx: &mut Self::Context) {
        ctx.binary(n.0.to_string());
    }
}

impl Agent {
    fn is_initialized(&mut self) -> bool {
        // self.seq == 0
        self.agent_id.is_some()
    }
    fn on_init(&mut self, ctx: &mut <Self as Actor>::Context) {
        let id = self.agent_id.as_ref().unwrap().clone();
        self.app_store.do_send(Event::AddAgent(id, ctx.address()));
    }
    fn on_stop(&mut self, _ctx: &mut <Self as Actor>::Context) {
        let id = self.agent_id.as_ref().unwrap().clone();
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
                println!("env ID={} ~/kubotnetes/k0s.io miniclient2 :8000", id);
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
    store: web::Data<Addr<AppStore>>,
) -> Result<HttpResponse, Error> {
    let agent_id = agent_id.id.clone();
    let store = store.get_ref();
    let handler = Agent::new(agent_id, store.clone());

    store.do_send(Event::Inc);

    // ws::start(handler, &req, stream)

    let mut res = ws::handshake(&req)?;
    Ok(res.streaming(ws::WebsocketContext::with_codec(
        handler,
        stream,
        Codec::new().max_size(10 * 1024 * 1024), // 10mb frame limit
    )))
}

#[derive(Debug)]
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
        self.receive_from_client(msg.0, ctx);
    }
}

#[derive(Message)]
#[rtype(result = "()")]
struct SetClientAddr(Addr<Client>);

impl Handler<SetClientAddr> for Terminal {
    type Result = ();
    fn handle(&mut self, s: SetClientAddr, _ctx: &mut Self::Context) {
        println!("setting self.client_addr");
        self.client_addr = Some(s.0);
    }
}

impl Actor for Terminal {
    type Context = ws::WebsocketContext<Self>;
    fn started(&mut self, ctx: &mut Self::Context) {
        let id = self.agent_id.clone();

        self.app_store
            .do_send(Event::PairPendingClient(id, ctx.address()));

        println!("Terminal started");
    }
    fn stopped(&mut self, _ctx: &mut Self::Context) {
        // close client too
        println!("Terminal stopped -> Client.stop()");
        if let Some(client_addr) = self.client_addr.clone() {
            client_addr.do_send(Stop);
        }
    }
}

impl Terminal {
    fn forward_to_client(&mut self, msg: ws::Message, _ctx: &mut <Self as Actor>::Context) {
        if let Some(client_addr) = self.client_addr.clone() {
            client_addr.do_send(WsMessage(msg))
        }
    }
    fn receive_from_client(&mut self, msg: ws::Message, ctx: &mut <Self as Actor>::Context) {
        // let mut message = "".to_string();
        match msg {
            ws::Message::Text(m) => {
                // message = m.to_string();
                println!(
                    "terminal::receive_from_client::ws::Message::Text {}",
                    m.len()
                );
                ctx.text(m);
            }
            ws::Message::Binary(m) => {
                // let message = String::from_utf8_lossy(&m).to_string();
                // if m.len() != message.len() { println!("terminal::handle_msg_from_client::bin({}) != str({})", m.len(), message.len()); }
                println!("terminal::receive_from_client::bin({})", m.len());
                ctx.binary(m);
            }
            _ => (),
        };
        // println!("terminal::handle_msg({}): {}", message.len(), message);
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for Terminal {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        match msg {
            Ok(m) => {
                self.forward_to_client(m, ctx);
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
    let agent_id = agent_id.id.clone().unwrap();
    let client_addr: Option<Addr<Client>> = None;
    let store = store.get_ref();
    let handler = Terminal::new(agent_id.clone(), client_addr, store.clone());

    println!("/api/terminal?id={}", agent_id);

    // ws::start(handler, &req, stream)

    let mut res = ws::handshake(&req)?;
    Ok(res.streaming(ws::WebsocketContext::with_codec(
        handler,
        stream,
        Codec::new().max_size(10 * 1024 * 1024), // 10mb frame limit
    )))
}

#[get("/api/terminalv2")]
async fn api_terminal_v2(
    req: HttpRequest,
    stream: web::Payload,
    agent_id: web::Query<ID>,
    store: web::Data<Addr<AppStore>>,
) -> Result<HttpResponse, Error> {
    let agent_id = agent_id.id.clone().unwrap();
    let client_addr: Option<Addr<Client>> = None;
    let store = store.get_ref();
    let handler = Terminal::new(agent_id.clone(), client_addr, store.clone());

    println!("/api/terminalv2?id={}", agent_id);

    // ws::start(handler, &req, stream)

    let mut res = ws::handshake(&req)?;
    Ok(res.streaming(ws::WebsocketContext::with_codec(
        handler,
        stream,
        Codec::new().max_size(10 * 1024 * 1024), // 10mb frame limit
    )))
}

#[derive(Debug)]
struct Client {
    buffer: VecDeque<ws::Message>,
    continuation_frame_buffer: BytesMut,
    agent_id: String,
    agent_addr: Addr<Agent>,
    app_store: Addr<AppStore>,
    terminal_addr: Option<Addr<Terminal>>,
    command: Command,
}

#[derive(Clone, Debug)]
enum Command {
    Terminal,
    TerminalV2,
}

impl Command {
    fn to_string(&self) -> String {
        match self {
            Command::Terminal => "TERMINAL\n".into(),
            Command::TerminalV2 => "TERMINALV2\n".into(),
        }
    }
}

impl Client {
    fn new(
        agent_id: String,
        agent_addr: Addr<Agent>,
        app_store: Addr<AppStore>,
        command: Command,
    ) -> Self {
        Client {
            buffer: VecDeque::<ws::Message>::new(),
            continuation_frame_buffer: BytesMut::new(),
            agent_id,
            agent_addr,
            app_store,
            terminal_addr: None,
            command,
        }
    }
}

#[derive(Message)]
#[rtype(result = "()")]
struct SetTerminalAddr(Addr<Terminal>);

impl Handler<SetTerminalAddr> for Client {
    type Result = ();
    fn handle(&mut self, s: SetTerminalAddr, _ctx: &mut Self::Context) {
        println!("setting self.terminal_addr");
        self.terminal_addr = Some(s.0);
    }
}

impl Handler<WsMessage> for Client {
    type Result = ();
    fn handle(&mut self, msg: WsMessage, ctx: &mut Self::Context) {
        self.receive_from_terminal(msg.0, ctx);
    }
}

impl Actor for Client {
    type Context = ws::WebsocketContext<Self>;
    fn started(&mut self, ctx: &mut Self::Context) {
        ctx.set_mailbox_capacity(1024);
        let id = self.agent_id.clone();
        self.app_store
            .do_send(Event::AddPendingClient(id, ctx.address()));
        self.agent_addr.do_send(NewTunnel(self.command.to_owned()));
        println!("Client started");
    }
    fn stopped(&mut self, _ctx: &mut Self::Context) {
        println!("Client stopped");
    }
}

impl Client {
    fn forward_to_terminal(&mut self, msg: ws::Message, _ctx: &mut <Self as Actor>::Context) {
        // println!("client forward");
        // let terminal_addr = self.terminal_addr.clone().unwrap();
        if self.terminal_addr.is_none() {
            self.buffer.push_back(msg);
            println!("push_back!");
        } else {
            while let Some(m) = self.buffer.pop_front() {
                self.terminal_addr.as_ref().unwrap().do_send(WsMessage(m));
            }
            self.terminal_addr.as_ref().unwrap().do_send(WsMessage(msg));
        }
    }
    fn receive_from_terminal(&mut self, msg: ws::Message, ctx: &mut <Self as Actor>::Context) {
        match msg {
            ws::Message::Text(m) => {
                ctx.text(m);
            }
            ws::Message::Nop => (),
            ws::Message::Ping(_) => (),
            ws::Message::Pong(_) => (),
            ws::Message::Close(_) => (),
            ws::Message::Binary(m) => {
                ctx.binary(m);
            }
            ws::Message::Continuation(c) => match c {
                Item::FirstText(ref m) | Item::FirstBinary(ref m) => {
                    self.continuation_frame_buffer.clear();
                    self.continuation_frame_buffer.put(m.to_owned());
                }
                Item::Continue(ref m) => {
                    self.continuation_frame_buffer.put(m.to_owned());
                }
                Item::Last(ref m) => {
                    self.continuation_frame_buffer.put(m.to_owned());
                    ctx.binary(self.continuation_frame_buffer.clone());
                    self.continuation_frame_buffer.clear();
                }
            },
        };
        // println!("client::handle_msg({}): {}", message.len(), message);
    }
}

impl StreamHandler<Result<ws::Message, ws::ProtocolError>> for Client {
    fn handle(&mut self, msg: Result<ws::Message, ws::ProtocolError>, ctx: &mut Self::Context) {
        // self.ensure_terminal();
        match msg {
            Ok(m) => {
                self.forward_to_terminal(m, ctx);
            }
            _ => {}
        }
    }
}

#[get("/api/agent/{id}/terminal")]
async fn api_client_terminal(
    req: HttpRequest,
    stream: web::Payload,
    id: web::Path<ID>,
    store: web::Data<Addr<AppStore>>,
) -> Result<HttpResponse, Error> /*impl Responder */ {
    println!("path: {}\n", req.path());

    let agent_id = id.id.clone().unwrap();
    let (contains, agent_addr) = store.send(ContainsAgent(agent_id.clone())).await.unwrap();
    if !contains {
        return Ok(HttpResponse::NotFound().body("not found"));
    }
    let store = store.get_ref();
    let handler = Client::new(
        agent_id,
        agent_addr.unwrap().clone(),
        store.clone(),
        Command::Terminal,
    );
    // ws::start(handler, &req, stream)

    let mut res = ws::handshake(&req)?;
    Ok(res.streaming(ws::WebsocketContext::with_codec(
        handler,
        stream,
        Codec::new().max_size(10 * 1024 * 1024), // 10mb frame limit
    )))
}

#[get("/api/agent/{id}/terminalv2")]
async fn api_client_terminal_v2(
    req: HttpRequest,
    stream: web::Payload,
    id: web::Path<ID>,
    store: web::Data<Addr<AppStore>>,
) -> Result<HttpResponse, Error> /*impl Responder */ {
    println!("path: {}\n", req.path());

    let agent_id = id.id.clone().unwrap();
    let (contains, agent_addr) = store.send(ContainsAgent(agent_id.clone())).await.unwrap();
    if !contains {
        return Ok(HttpResponse::NotFound().body("not found"));
    }
    let store = store.get_ref();
    let handler = Client::new(
        agent_id,
        agent_addr.unwrap().clone(),
        store.clone(),
        Command::TerminalV2,
    );
    // ws::start(handler, &req, stream)

    let mut res = ws::handshake(&req)?;
    Ok(res.streaming(ws::WebsocketContext::with_codec(
        handler,
        stream,
        Codec::new().max_size(10 * 1024 * 1024), // 10mb frame limit
    )))
}

#[derive(Deserialize)]
struct ID {
    id: Option<String>,
}

#[derive(Clone, Debug)]
struct AppStore {
    counter: usize,
    agents: HashMap<String, Addr<Agent>>,
    pending_clients: HashMap<String, Addr<Client>>,
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
        }
    }
}

#[derive(Message)]
#[rtype(result = "()")]
enum Event {
    Inc,
    AddAgent(String, Addr<Agent>),
    RemoveAgent(String),
    AddPendingClient(String, Addr<Client>),
    PairPendingClient(String, Addr<Terminal>),
}

impl Handler<Event> for AppStore {
    type Result = ();
    fn handle(&mut self, e: Event, _ctx: &mut Self::Context) {
        match e {
            Event::Inc => self.inc(),
            Event::AddAgent(id, addr) => self.add_agent(id, addr),
            Event::RemoveAgent(id) => self.remove_agent(id),
            Event::AddPendingClient(id, addr) => self.add_pending_client(id, addr),
            Event::PairPendingClient(id, addr) => self.pair_pending_client(id, addr),
        }
    }
}

#[derive(Message)]
#[rtype(result = "(bool, Option<Addr<Agent>>)")]
struct ContainsAgent(String);

impl Handler<ContainsAgent> for AppStore {
    type Result = MessageResult<ContainsAgent>;
    fn handle(&mut self, c: ContainsAgent, _ctx: &mut Self::Context) -> Self::Result {
        MessageResult(self.has_agent(c.0))
    }
}

#[derive(Message)]
#[rtype(result = "AppStore")]
struct AppState;

impl Handler<AppState> for AppStore {
    type Result = MessageResult<AppState>;
    fn handle(&mut self, _a: AppState, _ctx: &mut Self::Context) -> Self::Result {
        MessageResult(self.clone())
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
    fn add_pending_client(&mut self, agent_id: String, client_addr: Addr<Client>) {
        *self
            .pending_clients
            .entry(agent_id.clone())
            .or_insert(client_addr.clone()) = client_addr.clone();
        println!("store pending clients added: {}", agent_id);
    }
    fn pair_pending_client(&mut self, agent_id: String, terminal_addr: Addr<Terminal>) {
        let client_addr: Addr<Client>;

        {
            client_addr = self.pending_clients.get(&agent_id.clone()).unwrap().clone();
        }

        {
            self.pending_clients.remove(&agent_id.clone());
        }
        println!("store pending clients removed: {}", agent_id);

        client_addr.do_send(SetTerminalAddr(terminal_addr.clone()));
        terminal_addr.do_send(SetClientAddr(client_addr.clone()));

        println!(
            "client terminal pairs added: {:?} <=> {:?}",
            client_addr, terminal_addr
        );
    }
}

#[get("/state")]
async fn state(store: web::Data<Addr<AppStore>>) -> impl Responder {
    let state = store.send(AppState).await.unwrap();
    format!("{:?}", state)
}

#[get("/inc")]
async fn inc(store: web::Data<Addr<AppStore>>) -> impl Responder {
    store.do_send(Event::Inc);
    format!("{}", "Ok")
}

/// Telemetry entry point. Listening by default on 127.0.0.1:8000.
/// This can be changed using the `PORT` and `BIND` ENV variables.
#[actix_rt::main]
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
            .service(state)
            .service(api_client_terminal)
            .service(api_client_terminal_v2)
            .service(api_terminal)
            .service(api_terminal_v2)
            .service(api_agent)
    })
    .bind("127.0.0.1:8000")?
    .run()
    .await
}
