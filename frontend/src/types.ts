// src/types.ts
export interface Item {
    category: string;
    max: number;
    target: number;
    bad: number;
    good: number;
    actual: number;
   }
   
   export interface MonthData {
    month: string;
    max: number;
    actual: number;
    target: number;
    bad: number;
    good: number;
    items: Item[];
   }
   
   export interface SalesData {
    month: MonthData[];
    category: Item[];
   }
   