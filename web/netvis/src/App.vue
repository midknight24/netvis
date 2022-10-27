<script>
  import flowTable from './components/flowTables.vue'
  import flowMap from './components/flowMap.vue'
  export default {
    data() {
      return {}
    },
    components: {flowTable, flowMap}
  }
</script>

<template>
<div class="layout">
  <el-container>
    <el-main width="70%"><flowMap></flowMap></el-main>
    <el-aside width="30%"><flowTable></flowTable></el-aside>
  </el-container>
</div>
</template>

<style scoped>
</style>
