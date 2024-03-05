<template>
  <div class="container">
    <div class="main-chart">
      <BarChart :chartHeight="'500px'" />
    </div>

    <div class="vertical-graphs">
      <DgChart />
    </div>
  </div>
</template>

<script>
import BarChart from './components/Bar.vue';
import DgChart from './components/Doughnut.vue'

export default {
  components: {
    BarChart,
    DgChart
  },
};
</script>

<style scoped>
.container {
  display: flex;
  max-height: 100vh;
  overflow: hidden;
  padding: 5px;
}

.main-chart {
  width: 70%;
  height: 100vh;
  padding: 10px;
  justify-content: center;
  align-items: center;
  display: flex;
}

.vertical-graphs {
  width: 30%;
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 5px;
}
</style>
