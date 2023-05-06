<template>
  <div class="hi">
    Hi Vue ! {{ username }}
    <input type="button" @click="say_hi('horse')" value="Show" /> <br/><br/>

    <p>Vue.js周边的技术生态有： </p>
    <br/>
    <ul>
      <template v-for="tech in technologies" v-bind:key="tech">
        <li v-if="tech.indexOf('n') === 0">
          {{ tech }}
        </li>
      </template>
    </ul>

    <p>我们要使用的技术是： </p>

    <div v-if="name === 'Vue.js'">
      Vue.js !
    </div>
    <div v-else-if="name === 'angular'">
      Angular
    </div>
    <div v-else>
      React
    </div>

    <p v-bind:style="'color:' + my_color">Vue.js 学起来好开心~ (颜色： {{my_color}}) </p>
    <button @click="change_color()">变颜色</button>

    {{ message }} <button v-on:click='highlight' style='margin-top: 50px'>真的吗</button>


<!--    <p> 你好， {{ name }} ! </p>-->
<!--    <input type='text' v-model="name" />-->

    <div>
      <p>页面上的值: {{ name }} </p>
      <p> 通过视图层，修改my_value: </p>
      <input v-model="name" style='width: 400px'/>

      <hr/>
      <input type='button' @click="change_my_name_by_code()" value='通过控制代码修改my_value'/>
      <hr/>
      <input type='button' @click="show_my_name()" value='显示代码中的my_value'/>
    </div>


  </div>




</template>

<script>
export default {
  name: "SayHi",
  data() {
    return {
      // 原书：这里定义了 app, 就可以在其他地方通过 app.$data.name 来修改name 变量
      name: "Vue.js",
      username: "weilinux",
      my_color: 'green',
      message: '学习Vue.js使我快乐~ ',
      technologies: [
        "nvm", "npm", "node", "webpack", "ecma_script", "num2", "webpack2"
      ]
    }
  },
  methods: {
    say_hi: function(name) {
      alert('hi, ' + name + this.username)
    },
    change_color() {
      this.my_color = this.my_color == 'green' ? 'red' : 'green'
      console.info("== 新颜色是: ", this.my_color)
    },
    highlight: function() {
      this.message = this.message + '是的!'
    },
    show_my_name: function() {
      alert('my_name'+ this.name)
    },
    change_my_name_by_code: function() {
      this.name += ",  在代码中做修改, 666."
    }


  }

}
</script>

<style scoped>
.hi {
  color: red;
  font-size: 20px;
  font-style: italic;
}

</style>