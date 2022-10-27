<template>
  <div ref="flowmap" class="graph"></div>
</template>

<script>
import * as echarts from 'echarts';
import worldJson from '../assets/world.json'
export default {
    data() {
        return {
            flows: [],
            currentLocation: [0, 0]
        }
    },
    computed: {
        options() {
            return {
                backgroundColor: "#404a59",
                title: {
                    text: 'Netvis',
                    left: 'center',
                    textStyle: {
                        color: "#fff"
                    }
                },
                geo: {
                    zoom: 1.2,
                    type: "map",
                    map: "world",
                    roam: true,
                    label: {
                        emphasis: {
                            color: "#fff"
                        }
                    },
                    itemStyle: {
                        normal: {
                            areaColor: "#323c48",
                            borderColor: "#404a59"
                        },
                        emphasis: {
                            areaColor: "#2a333d"
                        }
                    }
                },
                series: [{
                    name: "你的位置",
                    type: "effectScatter",
                    coordinateSystem: "geo",
                    data: [this.currentLocation]
                }]
            }
        }
    },
    watch: {
        options() {
            this.chart.setOption(this.options)
        }
    },
    mounted() {
        this.init()
        navigator.geolocation.getCurrentPosition(location=>{
            this.currentLocation = [location.coords.longitude,location.coords.latitude]
        })
    },
    methods: {
        init() {
            if (this.chart){ return ;}
            let chart = echarts.init(this.$refs.flowmap);
            echarts.registerMap("world", worldJson)
            chart.setOption(this.options)
            this.chart = chart
        }
    }
}
</script>

<style>
    .graph {
    height: 550px;
    width: 100%;
    }
</style>