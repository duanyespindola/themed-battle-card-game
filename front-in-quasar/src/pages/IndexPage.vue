<script setup lang="ts">
import { ref, watch } from 'vue'
import { useStore } from 'src/stores/store';
import { connect, disconnect, sendMessage } from 'src/services/websocket';

const store = useStore()

const toggle_connect = ref(false)
watch(toggle_connect, (newVal: boolean)=>{
  if ( newVal && store.host_status === 'Not Connected' ){
    store.host_status = 'Connecting '
    connect()
  } else if (!newVal && store.host_status === 'Connected' ){ 
    disconnect()
  }
})
watch( ()=> store.host_status, (newVal: string)=>{
  if ( newVal === 'Connected' ){
    toggle_connect.value = true
  } else if (newVal === 'Not Connected' ){ 
    toggle_connect.value = false
  }
})


// const messages = store.messages
const message = ref('')

</script>


<template>
  <q-page class="items-center ">
    <q-toolbar class="bg-primary text-white">
      <q-toolbar-title>Send a message to echo server</q-toolbar-title>
      <div class="col-6">
        <q-input
          v-model="message"
          label="Message"
          outlined
          dense
          clearable
          class="bg-white"
          :disable="store.host_status !== 'Connected'">
          <template v-slot:append>
            <q-btn size="xs" color="primary"  icon="send" @click="sendMessage(message)" />
          </template>          
        </q-input>  
        
      </div>
    </q-toolbar>
    <div class="column">
      <div class="bg-red self-end q-pr-sm">
          <q-toggle
            v-model="toggle_connect"
            :label="store.host_status"
          />
      </div>
      <div class="row justify-center">
        <div class="col col-5 bg-blue">
          <q-chat-message
            v-for="msg in store.messages"
            :key="String(msg[0])"
            :text="[String(msg[2])]"
            :sent="msg[1] === 'Me'"
          />
        </div>

      </div>
    </div>

  </q-page>
</template>


