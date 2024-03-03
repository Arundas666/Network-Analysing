// src/data.ts
import type { MonthData, SalesData ,Item} from './types';

function getData(): MonthData[] {
 let year = '2018';
 const months = ['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'];
 const categories = ['Music', 'Video', 'Books', 'Electronics', 'Computers'];
 let data: MonthData[] = [];

 months.forEach(month => {
    let val: MonthData = {
      month: `${month} ${year}`,
      max: 50000,
      actual: 0,
      target: 25000,
      bad: 15000,
      good: 40000,
      items: []
    };

    categories.forEach(category => {
      let v: Item = {
        category: category,
        max: 10000,
        target: 5000,
        bad: 3000,
        good: 8000,
        actual: 0
      };
      let actual = Math.random();
      actual = Math.round(actual * 10000);
      v.actual = actual;
      val.actual += actual;
      val.items.push(v);
    });

    data.push(val);
 });

 return data;
}

export function getSales(): SalesData {
 let data = getData();
 let cats: Item[] = [];

 data.forEach(d => {
    let items = d.items;
    items.forEach((item, idx) => {
      if (idx >= cats.length) {
        cats.push({
          category: item.category,
          max: 0,
          target: 0,
          bad: 0,
          good: 0,
          actual: 0
        });
      }
      let cat = cats[idx];
      cat.max += item.max;
      cat.target += item.target;
      cat.bad += item.bad;
      cat.good += item.good;
      cat.actual += item.actual;
    });
 });

 return {
    month: data,
    category: cats
 };
}
