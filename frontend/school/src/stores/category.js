const { ref } = require("vue");


export const categoryDatas = defineStore('category',()=>{
    const categoryDatas = ref([])
    return {categoryDatas}
})