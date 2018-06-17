<template>
  <div>
    <div>
      <input v-model="message" type="text">
      <button v-on:click="submit">送信</button>
    </div>
    <div>
      <ul>
        <li v-for="content in contents">
          {{content}}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
  import axios from "axios"

  export default {
    name: 'BBS',
    data () {
      return {
        contents: ["読み込み中"]
      }
    },
    state: {
      message: ''
    },
    created: function() {
      const self = this;
      axios.get("http://localhost:8888/user_posts/")
        .then(function(response) {
          self.contents = [];
          
          response.data["messages"].forEach( message =>
            self.contents.unshift(message)
          );
        })
        .catch(function(error) {
          console.log('エラーだよ: ' + error)
        });
    },
    methods: {
      submit: function(event) {
        if (this.message === '') return;

        const self = this;
        const data = {message: this.message};
        axios.post("http://localhost:8888/user_posts/", data)
          .then(function() {
            self.contents.unshift(self.message);
            self.message = '';
          })
          .catch(function(error) {
            console.log('エラーだよ: ' + error)
          });
      }
    }
  }
</script>

<style>
  li {
    list-style: none;
  }
</style>
