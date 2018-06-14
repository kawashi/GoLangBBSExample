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
  export default {
    name: 'BBS',
    data () {
      return {
        contents: []
      }
    },
    state: {
      message: ''
    },
    methods: {
      submit: function(event) {
        // TODO: 送信クリック後の処理
        // TODO: 1. 画面に反映
        // TODO: 2. GoのAPIの/user_postsにPOST
        if (this.message === '') return;
        this.contents.unshift(this.message);
        this.message = '';
        vue.$forceUpdate();
      }
    },
    created: function() {
      // TODO: 初期化処理
      // TODO: 1. GoのAPIの/user_postsをGETして、contentsに反映
      this.contents.unshift("1件目の投稿");
      this.contents.unshift("2件目の投稿");
    }
  }
</script>

<style>
  li {
    list-style: none;
  }
</style>
