<template>
    <div class="container">
        <div v-if="!loaded">Loading...</div>
        <div v-else>
            <Bar :data="chartData" ref="chart" :style="{ height: chartHeight }" :options="chartOptions" />
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
                onClick: this.handleChartClick,
            },
            countryId: [],
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
                let barData = result.data[0]
                let labels = [];
                let speed = []

                barData.forEach((item) => {
                    labels.push(item.Country)
                    speed.push(item.OverallSpeed)
                    this.countryId.push(item.CountryId)
                })

                this.chartData = {
                    labels: [
                        ...labels
                    ],
                    datasets: [
                        {
                            label: 'Data 4',
                            backgroundColor: '#f87979',
                            data: [...speed]
                        }
                    ]
                };

                this.loaded = true;
            } catch (e) {
                console.error(e);
            }
        },
        handleChartClick(event, array) {
            if (array.length !== 0) {
                const position = array[0].index;
                let selectedCountryId = this.countryId[position]
                this.loadAdditionalData(selectedCountryId);
            } else {
                console.log("You selected the background!");
            }
        },
        async loadAdditionalData(countryId) {
            try {
                const response = await fetch(`http://localhost:8080/avg-in-states/${countryId}`);

                if (!response.ok) {
                    throw new Error(`HTTP error! Status: ${response.status}`);
                }

                const result = await response.json();
                let states = [];
                let speed = []
                result.data[0].forEach((item) => {
                    states.push(item.State)
                    speed.push(item.OverallSpeed)
                })
                const backgroundColors = this.generateRandomColors(states.length);
                const DoughnutData =
                {
                    labels: [...states],
                    datasets: [
                        {
                            backgroundColor: backgroundColors,
                            data: [...speed]
                        }
                    ]
                }
                this.$emit('doughnutDataLoaded', DoughnutData);
            } catch (e) {
                console.error(e);
            }
        },
        generateRandomColors(length) {
            const colors = [];
            for (let i = 0; i < length; i++) {
                const color = `#${Math.floor(Math.random() * 16777215).toString(16)}`;
                colors.push(color);
            }
            return colors;
        },
    }
}
</script>
