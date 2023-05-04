<template>
    <div class="">
        <form @submit.prevent="createMock">
            <div>
                <p>
                    <label>apiName</label>
                    <input type="text" v-model="mock.apiName" />
                </p>

                <p>
                    <label>key</label>
                    <input type="text" v-model="mock.reqKey" />
                </p>
                <p>
                    <select v-model="valueType">
                        <option>text</option>
                        <option>number</option>
                        <option>select</option>
                    </select>
                </p>
                <p v-if="valueType=='text'">
                    <label>val</label>
                    <input type="text" v-model="mock.reqValue" />
                </p>
                <p v-if="valueType=='number'">
                    <label>val</label>
                    <input type="number" v-model="mock.reqValue" />
                </p>
                <div v-if="valueType=='select'">
                    <label>val</label>
                    <h4>select</h4>
                </div>
                <p>
                    <v-jsoneditor
                        ref="jsonEditor"
                        class="col-5 p-0"
                        v-model="mock.reply"
                        :options="{
              mode: 'code',
              mainMenuBar: false,
              navigationBar: false,
            }"
                    ></v-jsoneditor>
                </p>
                <p>
                    <button >create</button>
                </p>
            </div>
        </form>
        <div>
            <h1>Preview:</h1>
                <Card :mock="mock" />
        </div>
    </div>
</template>

<script>

import Card from "@/components/Card.vue";
import VJsoneditor from 'v-jsoneditor'

export default {
    name: 'CreateMock',
    components: {
        Card,
        VJsoneditor
    },
    data(){
        return{
            valueType:"text",
            mock:{
                "id":"",
                "apiName":"",
                "reqKey":"",
                "value":null,
                "handlerType":"api",
                "reply":{},
                "timeOut":0
            }
        }
    },
    computed:{
    },
    methods:{
        createMock(){
            this.mock.reply=JSON.stringify(this.mock.reply)
            console.log(this.mock)
            console.log(JSON.stringify(this.mock.reply))
            this.$store.dispatch('createMock', this.mock).then(() => {
                console.log("create")
                this.$store.dispatch('getAllMocks')
                this.$router.push('/')
            })
                .catch(err => console.log(err))
        }
    }
}
</script>
