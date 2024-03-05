<template>
    <div class="container">
        <div v-if="!loaded">Loading...</div>
        <div v-else>
            <Bar :data="chartData" :style="{ height: chartHeight }" :options="chartOptions" />
        </div>
    </div>

</template>

<script>
import { Bar } from 'vue-chartjs'
import { Chart as ChartJS, Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale } from 'chart.js'

ChartJS.register(Title, Tooltip, Legend, BarElement, CategoryScale, LinearScale)

export default {
    name: 'BarChart',
    components: { Bar },
    props: {
        chartHeight: {
            type: String,
        },
    },
    data() {
        return {
            loaded: false,
            chartData: null,
            chartOptions: {
                responsive: true,
                maintainAspectRatio: false,
            },
        }
    },
    mounted() {
        this.loadData();
    },
    methods: {
        async loadData() {
            this.loaded = false;
            try {
                const response = await fetch('http://localhost:8080/find-avg');

                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }

                const result = await response.json();
                console.log(result)
            const a=result.data[0]
            let labels =[]
            let data=[]
                a.forEach((item)=>{
                    labels.push(item.Country)
                    data.push(item.OverallSpeed)
                })

                this.chartData = {
                    labels: [
                       ...labels
                    ],
                    datasets: [
                        {
                            label: 'Average Network speed',
                            backgroundColor: '#f87979',
                            data: [...data]
                        }
                    ]
                };

                this.loaded = true;
            } catch (e) {
                console.error(e);
            }
        }
    }
}
</script>
