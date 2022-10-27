import { defineStore } from 'pinia'
import axios from 'axios'

export const useFlowStore = defineStore('flow', {
  state: () => {
    return {
        flows: []
    }
  },
  actions: {
    async getFlows() {
        try {
            const data = await axios.get('http://127.0.0.1:80')
            console.log(data)
            this.flows = data
        } catch(error) {
            alert(error)
            console.log(error)
        }
    }
  },
})