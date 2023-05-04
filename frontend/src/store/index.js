import Vue from 'vue'
import Vuex from 'vuex'
import axios from "axios";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    mocks:[]
  },
  mutations: {
    SET_MOCKS(state, payload) {
      Vue.set(state, 'mocks', payload)
    },
  },
  actions: {
    getAllMocks({commit}) {
      return axios.post("http://127.0.0.1:45765/getAllMocks",{})
          .then(res=>{
            console.log(res.data)
            commit('SET_MOCKS',res.data.data)
          })
          .catch(err => console.log(err))
    },
  },
  modules: {
  },
  getters:{
    getAllMocks: (state) => {
      return state.mocks
    },
  }
})
