<template>
    <div class="container-fluid">
     
       <div class="form-group">
            <div class="row">
                <wj-flex-chart header="2018 Annual Sales" bindingX="month" selectionMode="Point" :initialized="initializeChart"
                        :itemsSource="chartData" :selectionChanged="chartSelectionChanged">
                    <wj-flex-chart-legend position="None"></wj-flex-chart-legend>
                    <wj-flex-chart-series binding="actual" name="Sales"></wj-flex-chart-series>
                </wj-flex-chart>
            </div>
            <div class="row">
                <div class="col">
                    <wj-flex-pie :header="pieHeader" bindingName="category" binding="actual" :itemsSource="pieData" 
                            :labelContent="pieLabel" :initialized="initializePie">
                        <wj-flex-chart-legend position="None"></wj-flex-chart-legend>
                    </wj-flex-pie>
                </div>
                <div class="col">
                    <ul class="bullets">
                        <li v-for="item in bulletsData">
                            <label>{{item.category}}</label>
                            <wj-bullet-graph showText="Value" :target="item.target" :max="item.max" :good="item.good" :bad="item.bad" :value="item.actual">
                            </wj-bullet-graph>
                        </li>
                    </ul>
                </div>
            </div>
        </div>
    </div>
   </template>
   
   <script lang="ts">
   import { defineComponent, ref } from 'vue';
   import { getSales } from './data';
   import { SalesData } from './types';
   
   export default defineComponent({
    name: 'App',
    setup() {
       const data = ref<SalesData>(getSales());
       const chartData = ref(data.value.month);
       const pieData = ref(data.value.category);
       const bulletsData = ref(data.value.category);
       const pieHeader = ref('2018 Product Sales');
   
       // Define methods and lifecycle hooks here using the Composition API
   
       return {
         data,
         chartData,
         pieData,
         bulletsData,
         pieHeader,
         // Return other methods and lifecycle hooks
       };
    },
   });
   </script>
   
   <style>
    body {
        margin-bottom: 24px;
        min-width: 40em;
        overflow: auto;
    }

    .wj-flexchart {
        border: none;
    }

    .col {
        float:left;
        width: 50%;
    }

    li {
        list-style-type: none;
    }

    .wj-gauge {
        margin: 17px 1em;
    }

    .wj-gauge svg {
        overflow: visible !important; 
    }
</style>
   