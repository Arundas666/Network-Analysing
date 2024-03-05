<template>
    <div class="container">
        <div v-if="!loaded">Loading...</div>
        <div v-else>
            <Doughnut :data="chartData" :style="{ height: chartHeight }" :options="chartOptions" />
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
    data() {
        return {
            loaded: false,
            chartData: null,
            chartOptions: {
                responsive: true,
                maintainAspectRatio: false
            }
        }
    },
    mounted() {
        this.loadData();
    },
    methods: {
        async loadData() {
            this.loaded = false;
            try {
                const response = await fetch('http://localhost:8080/avg-in-states');

                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }

                const result = await response.json();

                this.chartData = {
                    labels: ['apple', 'orange', 'kappikuru', 'andiputtu'],
                    datasets: [
                        {
                            backgroundColor: ['#41B883', '#E46651', '#00D8FF', '#DD1B16'],
                            data: [40, 20, 80, 10]
                        }
                    ]
                }

                this.loaded = true;
            } catch (e) {
                console.error(e);
            }
        }
    }
}
</script>
