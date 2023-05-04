<template>
    <div class="warp">
        <div class="cardwarp">
            <div class="card">
                <div class="card-header">/{{ mock.apiName}}</div>
                <div class="card-body">
                        <h3>Req:</h3>
                        <pre>{"{{mock.reqKey}}":{{mock.reqValue}}}</pre>
                        <h3>Res:</h3>
                        <pre>{{ mock.reply}}</pre>
                </div>
                <div class="card-footer">
                    <button class="card-btn">edit</button>
                    <button class="card-btn" @click="deleteById">delete</button>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
export default {
    name: 'Card',
    props:["mock"],
    methods:{
        deleteById(){
            if (confirm("Are you sure?")){
                this.$store.dispatch('deleteMockById', {id: this.mock.id}).then(() => {
                    console.log("delete")
                })
                    .catch(err => console.log(err))
            }
        }
  }
}
</script>

<style scoped>

.card-btn{
    border-radius: 1rem;
    background-color: rgb(0, 69, 138);
    border: 0.1rem solid transparent;
    color: rgb(255, 255, 255);
    font-size: 1rem;
    width: 5rem;
    height: 2.2rem;
    cursor: pointer;
    font-family: OpenSans-Bold, HelveticaNeue-Bold;
}
.card-footer{
    display: flex;
    justify-content: space-between;
    border-top: 0.1rem solid rgb(205, 205, 205);
    padding: 0.5rem;
    margin-right: 1rem;
    margin-left: 1rem;
}
.card-header{
    display: flex;
    justify-content: space-between;
    border-bottom: 0.1rem solid rgb(205, 205, 205);
    padding: 0.5rem;
    margin-right: 1rem;
    margin-left: 1rem;
    font-size: 1.5rem;
}
.card-body{
    display: flex;
    align-items: center;
    padding: 1rem;
    gap: 2.5rem;
}

.cardwarp{
    display: flex;
    justify-content: space-evenly;
    align-items: center;
}
.card {
    display: flex;
    flex-direction: column;
    background-color: rgb(255, 255, 255);
    justify-content: center;
    border-radius: 0.625rem;
    cursor: pointer;
    height: 12.8rem;
    width: 40rem;
    border: 4px solid rgba(0, 69, 138, 0.5);
    box-shadow: rgb(0 0 0 / 55%)
}

.warp {
    display: flex;
    flex-direction: column;
    gap: 1rem;
    max-height: 100%;
    overflow: auto;
    padding: 1rem;
}


@media only screen and (max-width: 600px) {
    .cardwarp{
        width: 100%;
        display: flex;
        flex-direction: column;
        gap: 1rem;
        overflow: auto;
        margin-top: 2rem;
        padding-bottom: 2.5rem;
        height: 758.304px;    }
    .card{
        display: flex;
        flex-direction: column;
        background-color: rgb(255, 255, 255);
        justify-content: center;
        max-width: 55rem;
        width: 90%;
        border-radius: 0.625rem;
        min-height: 11rem;
        overflow: hidden;
        border: 0.1rem solid rgb(205, 205, 205);
        box-shadow: rgb(0 0 0 / 15%) 0px 2px 10px 0px;    }
}

</style>
