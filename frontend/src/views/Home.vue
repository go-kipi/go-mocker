<template>
  <div class="home">
      <div >
          <Card v-for="mock in mocks" :key="mock.id" :mock="mock" />
      </div>
  </div>
</template>

<script>

import Card from "@/components/Card.vue"
export default {
  name: 'Home',
  components: {
      Card,
  },
    data(){
      return{
      }
    },
    computed:{
        mocks() {
            return this.$store.getters.getAllMocks
        },
    },
    methods:{

    }
}
</script>
