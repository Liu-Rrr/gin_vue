<template>
  <div style="width: 100%; height: 100vh ; background: #2c3e50 ; overflow: hidden">
    <div style="width: 350px; margin: 100px auto">
      <div style="color: #cccccc;font-size: 30px ; text-align: center ; padding: 30px 0">注册</div>
      <el-form ref="form" :model="form" :rules="rules" size="small" >
        <el-form-item prop="username">

          <el-input prefix-icon="el-icon-user" v-model="form.username" size="medium" clearable></el-input>

        </el-form-item>

        <el-form-item prop="password">

          <el-input prefix-icon="el-icon-key" v-model="form.password" size="medium" show-password></el-input>

        </el-form-item>

        <el-form-item prop = "confirm">

          <el-input prefix-icon="el-icon-key" v-model="form.confirm" size="medium" show-password></el-input>

        </el-form-item>

        <el-form-item>
          <el-button style="width: 100% " type="primary" size="medium" @click="register">注册</el-button>
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
  name: "Register.vue",
  data(){
    return{
      form:{},
      rules:{
        username:[
          { required: true, message: '请输入姓名', trigger: 'blur' },
        ],
        password:[
          { required: true, message: '请输入密码', trigger: 'blur' },
        ],
        confirm:[
          { required: true, message: '请重复密码', trigger: 'blur' },
        ]
      }
    }
  },

  methods:{
    register() {
      if (this.form.confirm !== this.form.password) {
        this.$message({
          type: "success",
          message: "密码不一致"
        })
      } else {
        this.$refs['form'].validate((valid) => {
          if (valid) {
            let form = (({username, password}) => ({username, password}))(this.form);
            console.log(form)
            request.post("/user/Register", form).then(res => {
              if (res.code === '0') {
                this.$message({
                  type: "success",
                  message: "登录成功"
                })
                // this.load()
                this.$router.push("/login")//element-plus登录成功后页面跳转
              } else {
                this.$message({
                  type: "error",
                  message: res.msg
                })
              }
            })
          }
        })

      }
    }
  }
}
</script>

<style scoped>

</style>