<template>
  <div class="hello">
    <h3>Pusher Test</h3>
    <div class="message-box">
      <div class="input-box">
        <input type="text" id="message" name="message" v-model="message">
        <button @click="sendMessage">send Message</button>
      </div>
      <div class="container-message">
        <p v-for="(item, key) in listMessage" v-bind:key=key>
          {{item}}
        </p>
      </div>
    </div>
  </div>
</template>

<script>
import Pusher from 'pusher-js'
import axios from 'axios'
export default {

  name: 'HelloWorld',
  created () {
    this.subscribe()
  },
  data () {
    return {
      listMessage: [],
      message:''
    }
  },
  methods: {
    subscribe () {
      let pusher = new Pusher('279a52598e5317c2f0f9', { cluster: 'ap1',channelAuthorization: {
        endpoint: "http://192.168.1.6:8888/pusher/auth",
        headers: {authorization: "Bearer " + "idToken"},
      }, })
      pusher.subscribe('private-my-channel')
      pusher.bind('my-event', data => {
        this.listMessage.push(pusher.connection.socket_id +": "+ data.message)
      })
    },
    async sendMessage(){
      await axios.post("http://192.168.1.6:8888/messages", {"message": this.message})
    }
  },
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
input {
  margin-right: 10px;
}
</style>
