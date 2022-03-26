<template>

<!--  此时的view作为全局的分节点访问，根据路由展示-->
  <div>
    <router-view />
  </div>


</template>

<style>

</style>


<script>

import {ElConfigProvider} from 'element-plus'
export default {
  name:"App",
  components:{

  },
  data(){
    return{
    }
  }
}
</script>
