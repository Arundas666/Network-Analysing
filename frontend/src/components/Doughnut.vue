<template>
    <div class="container">
        <div v-if="!loaded">Loading...</div>
        <div v-else>
            <Doughnut :data="chartData"  :options="chartOptions" />
        </div>
    </div>
</template>

<script>
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from 'chart.js'
import { Doughnut } from 'vue-chartjs'

ChartJS.register(ArcElement, Tooltip, Legend)

export default {
    name: 'DtChart',
    components: { Doughnut },
    props: {
        chartData: {
            type: Object,
            default: null,
        },
    },
    data() {
        return {
            loaded: false,
            chartOptions: {
                responsive: true,
                maintainAspectRatio: false
            }
        }
    },
    mounted() {
        this.loaded = true;
    },
}
</script>
