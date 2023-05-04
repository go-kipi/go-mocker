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
    DELETE_MOCK(state, payload) {
      // var newMocks = state.mocks.filter(mock=> mock.id !== payload.id)
      // Vue.set(state, 'mocks', newMocks)
    },
    CREATE_MOCK(state, payload) {
    var newMocks={
      ...state.mocks,
      payload
    }
      Vue.set(state, 'mocks', newMocks)

    },
  },
  actions: {
    getAllMocks({commit}) {
      return axios.post("http://127.0.0.1:45765/getAllMocks",{})
          .then(res=>{
            commit('SET_MOCKS',res.data.data)
          })
          .catch(err => console.log(err))
    },
    createMock({commit},payload) {
      console.log(payload)
      return axios.post("http://127.0.0.1:45765/createMock",payload)
          .then(res=>{
            commit('CREATE_MOCK',payload)
          })
          .catch(err => console.log(err))
    },
    deleteMockById({commit},payload) {
      return axios.post("http://127.0.0.1:45765/deleteMockById",payload)
          .then(res=>{
            commit('DELETE_MOCK',payload)
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
