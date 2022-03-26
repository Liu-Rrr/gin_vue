<template>
  <div>
  <!--    头部-->
  <Header />
  <!--主体-->
  <div style="display: flex">
    <!--侧边栏-->
    <Aside />
    <!--内容-->
    <el-config-provider :locale="locale">
      <!--home-->
      <router-view style="flex: 1"/>
      <!--router-view为其他页面的显示-->

    </el-config-provider>
  </div>
  </div>

</template>

<script>
import Header from "@/components/Header";
import Aside from "@/components/Aside";
import zhCn from 'element-plus/lib/locale/lang/zh-cn'


export default {
  name:"Layout",
  components:{
    Aside,
    Header
  },
  data(){
    return{
      locale:zhCn,
    }
  }
}
</script>


<style scoped>

</style>