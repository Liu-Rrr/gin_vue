<template>
  <div style="width: 100%; height: 100vh ; background: #2c3e50 ; overflow: hidden">
    <div style="width: 350px; margin: 100px auto">
      <div style="color: #cccccc;font-size: 30px ; text-align: center ; padding: 30px 0">登录</div>
      <el-form ref="form" :model="form" :rules="rules"  size="small" >
        <el-form-item prop="username">

          <el-input prefix-icon="el-icon-user" v-model="form.username" size="medium" clearable></el-input>

        </el-form-item>

        <el-form-item prop="password">

          <el-input prefix-icon="el-icon-key" v-model="form.password" size="medium" show-password></el-input>

          </el-form-item>

        <el-form-item>
          <el-button style="width: 100% " type="primary" size="medium" @click="login">登录</el-button>
          </el-form-item>



<!--        <button-->
      </el-form>

    </div>
  </div>
</template>

//form一定要被定义，收不到数据无法显示，form.username一定要不能.name
<script>
import request from "@/untils/request";

export default {
  name: "login.vue",
  data(){
      return{
        form:{},
        rules: {
          username: [
            {required: true, message: '请输入姓名', trigger: 'blur'},
          ],
          password: [
            {required: true, message: '请输入密码', trigger: 'blur'},
          ],
        }
    }
  },

  methods:{
    login(){
      //验证成功后才发出请求
      this.$refs['form'].validate((valid) => {
        if (valid) {
          request.post("/user/login",this.form).then(res=>{
            if(res.code === '0'){
              this.$message({
                type:"success",
                message:"登录成功"
              })
              sessionStorage.setItem("user",JSON.stringify(res.data))//登录后设置缓存
              // this.load()
              this.$router.push("/user")//element-plus登录成功后页面跳转
            } else{
              this.$message({
                type:"error",
                message:res.msg
              })
            }
          })
        }
      })

    }
  }
}
</script>

<style scoped>

</style>