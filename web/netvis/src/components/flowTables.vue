<template>
    <el-table :data="flows" style="width: 100%">
    <el-table-column prop="scrIP" label="scrIP" width="180" />
    <el-table-column prop="dstIP" label="dstIP" width="180" />
    </el-table>
</template>

<script>
import {useFlowStore} from '../stores/flows'
export default {
    data() {
        return {
            flows: [
                {
                    scrIP: "192.168.31.40",
                    dstIP: "8.8.4.4"
                }
            ]
        }
    },
    mounted() {
        const flows = useFlowStore()
        flows.getFlows()
    }
}
</script>
<style>
</style>