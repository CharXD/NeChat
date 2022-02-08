<template>
  <div class="bg">
    <div class="Register">
      <div class="card">
        <div class="cardHead">
          <img src="../assets/logo.png" alt="">
          <p>NeChat - Register</p>
        </div>

        <div class="cardBody">
          <el-input class="username_input" v-model="username" placeholder="Your nick name..."></el-input>
          <el-input class="password_input" v-model="password" show-password placeholder="Enter your password..."></el-input>
          <el-input class="confirm_password_input" v-model="confirm_password" show-password placeholder="Enter your password again..."></el-input>
          <el-button @click="register">Register</el-button>
          <el-link href="#/Login" type="info">login</el-link>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios"
import {Message} from 'element-ui';

export default {
  name: 'Register',
  data() {
    return {
      img: require('../assets/background.png'),
      username: null,
      password: null,
      confirm_password: null
    }
  },
  methods: {
    register() {
      // this.$router.replace('#/Register');
      if (this.username === null || this.password === null || this.confirm_password === null) {
        Message({
          message: '用户名或密码不能为空', type: 'warning', duration: 1500
        });
        return;
      }

      if (this.password !== this.confirm_password) {
        Message({
          message: '两次输入密码不一致', type: 'warning', duration: 1500
        });
        return;
      }

      const u_pattern = /^[a-zA-Z0-9_-]{4,16}$/;
      const p_pattern =/^.*(?=.{6,})(?=.*\d)(?=.*[A-Z])(?=.*[a-z])(?=.*[!@#$%^&*?.]).*$/;

      if (!u_pattern.test(this.username)) {
        Message({
          message: '用户名必须符合4到16位（字母，数字，下划线，减号）', type: 'warning', duration: 1500
        });
        return;
      }

      if (!p_pattern.test(this.password)) {
        Message({
          message: '密码最少6位，包括至少1个大写字母，1个小写字母，1个数字，1个特殊字符', type: 'warning', duration: 1500
        });
        return;
      }

      axios.post('http://127.0.0.1:25566/register', {
        UserName: this.username,
        Password: this.password
      })
          .then(function (response) {
            const data = response.data;
            if (data.code === 0) {
              Message({
                message: data.msg, type: 'success', duration: 1500
              });
            } else {
              Message({
                message: data.msg, type: 'error', duration: 1500
              });
            }
          })
          .catch(function (error) {
            Message({
              message: error, type: 'error', duration: 1500
            });
          })
    }
  }
}
</script>
<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.bg{
  position: absolute;
  top: 0;
  left: 0;
  height: 100%;
  width: 100%;
  background: url(../assets/background.png) no-repeat fixed center center;
  background-size:cover;
}

.card {
  border-radius: 5px;
  width: 390px;
  height: 500px;
  background-color: #ffffff;
  position: absolute;
  left: 0;
  top: 0;
  bottom: 0;
  right: 0;
  margin: auto 15% auto auto;
  box-shadow: 0 0 10px 5px #eeeeee;
}

.cardHead {
  position: absolute;
  left: 42px;
  width: 316px;
  height: 131px;
}

.cardHead img {
  position: absolute;
  left: 0;
  right: 0;
  margin: 10% auto;
}

.cardHead p {
  position: absolute;
  margin: 40% 76px;
  font-size: 20px;
}

.cardBody {
  position: absolute;
  left: 42px;
  bottom: 0;
  width: 316px;
  height: 330px;
}

.el-input {
  margin-top: 25px;
}

.el-button, .el-button:focus {
  margin-top: 30px;
  width: 316px;
  background: #374b4f;
  border-color: #374b4f;
  color: #ffffff;
}

.el-button:hover {
  background: #465f65;
  border-color: #465f65;
  color: #ffffff;
}

.el-link {
  margin-top: 15px;
}
</style>
