<template>
  <div class="home">
      <div >
          <Card v-for="mock in mocks" :key="mock.id" :mock="mock" />
      </div>
<!--      {{  mocks}}-->
<!--      <Card />-->
<!--      <Card />-->
<!--      <Card />-->
<!--      <Card />-->
  </div>
</template>

<script>

import Card from "@/components/Card.vue"
import axios from 'axios'

export default {
  name: 'Home',
  components: {
      Card,
  },
    data(){
      return{
          // mocks:null
      }
    },
    computed:{
        mocks() {
            return this.$store.getters.getAllMocks
        },
    },
    methods:{
      getAllMocks(){
          axios.post("http://127.0.0.1:45765/getAllMocks",{})
              .then(res=>{
                  console.log(res.data)
                  this.mocks=res.data.data
              })
              .catch(err => console.log(err))
      }
    },
    async created() {
        await this.$store.dispatch('getAllMocks')
    }
}
</script>
