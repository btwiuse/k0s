    net.Conn

asciitransport /\
client <-> server

---

client.Read => Frame{Output} <= server.Write client.Write => Frame{Input,Resize}
<= server.Read

---

client << Output << server client >> Input/Resize >> server

---

os.Stdout client << Output << server term.Read []byte os.Stdin client >> Input
>> server term.Write []byte term.winch client >> Resize >> server term.Resize
uint32
