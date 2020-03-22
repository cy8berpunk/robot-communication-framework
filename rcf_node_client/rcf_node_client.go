package rcf_node_client

import(
  "fmt"
  "net"
  "strconv"
  "strings"
  "bufio"
  "bytes"
  "robot-communication-framework/rcf_util"
)

// function to connect to tcp server (node) and returns connection
func connect_to_tcp_server(port int) net.Conn{
  conn, err := net.Dial("tcp4", ":"+strconv.Itoa(port))

  if err != nil {
    fmt.Println("an error occured: ")
    fmt.Println(err)
  }
  // don't forget to close connection
  return conn
}

// returns connection to node
func Node_open_conn(node_id int) net.Conn {
  var conn net.Conn = connect_to_tcp_server(node_id)

  return conn
}

func Node_close_conn(conn net.Conn) {
  conn.Write([]byte("end\n"))
  conn.Close()
}

// pushes data to topic stack
func Topic_publish_data(conn net.Conn, topic_name string, data string) {
  conn.Write([]byte(topic_name+"+"+data+"\n"))
}

// pulls x elements from topic topic stack
func Topic_pull_data(conn net.Conn, nelements int, topic_name string) []string {
  conn.Write([]byte(topic_name+"-"+strconv.Itoa(nelements) + "\n"))
  var elements []string
  rdata, _ := bufio.NewReader(conn).ReadString('\n')
  elements = strings.Split(rcf_util.Trim_suffix(rdata), ",")

  return elements
}

// pushes data to topic stack
func Topic_glob_publish_data(conn net.Conn, topic_name string, data map[string]string) {
  encoded_data := []byte(rcf_util.Glob_map_encode(data).Bytes())
  bsend := append([]byte(topic_name+"+"), encoded_data...)
  conn.Write(bsend)
}

// pulls x elements from topic topic stack
func Topic_glob_pull_data(conn net.Conn, nelements int, topic_name string) []map[string]string {
  conn.Write([]byte(topic_name+"-"+strconv.Itoa(nelements) + "\n"))
  elements := make([]map[string]string, 0)
  rdata := make([]byte, 512)
  n, err_handle := bufio.NewReader(conn).Read(rdata)
  rdata = rdata[:n]
  if err_handle != nil {
    fmt.Println("/[read] ", err_handle)
  }

  fmt.Println("Data: ", rdata)
	split_rdata := bytes.Split(rdata, []byte("\r"))
  fmt.Println("split data: ", split_rdata)

  for _, map_element := range split_rdata {
    b := bytes.NewBuffer(make([]byte,0,len(map_element)))
    b.Write(rdata)

    decodedMap := rcf_util.Glob_map_decode(b)
    elements = append(elements, decodedMap)
  }
  return elements
}

// waits continuously for incoming topic elements, enables topic data streaming before
func Topic_subscribe(conn net.Conn, topic_name string) <-chan string{
  conn.Write([]byte("$"+topic_name+"\n"))
  topic_listener := make(chan string)
  go func(topic_listener chan<- string){
    for {
      data, err := bufio.NewReader(conn).ReadString('\n')
      topic_listener <- data
      if err != nil {
        fmt.Println("conn closed")
        break
      }
    }
  }(topic_listener)
  return topic_listener
}

//  creates new action on node
func Topic_create(conn net.Conn, topic_name string) {
  conn.Write([]byte("+"+topic_name + "\n"))
}

//  executes action
func Action_exec(conn net.Conn, action_name string) {
  conn.Write([]byte("*"+action_name + "\n"))
}

// lists node's topics
func Topic_list(conn net.Conn) []string {
  conn.Write([]byte("list_topics\n"))
  data, _ := bufio.NewReader(conn).ReadString('\n')

  return strings.Split(rcf_util.Trim_suffix(data), ",")
}
