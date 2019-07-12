const URL_App_Devices_All = 'http://192.168.11.4:8002/v1/device/all'
const URL_App_Devices_Add = 'http://192.168.11.4:8002/v1/device/add/'
const URL_App_Devices_Remove = 'http://192.168.11.4:8002/v1/device/remove/'

// register the devices component
var appDevices = Vue.component('app-devices', {
    data: function () {
        return {
            items: [],            
        }
    },
    created () {
        this.fetchData()
    },
    methods: {
        fetchData () {
            var self = this;
            self.items = []
            this.$http.get(URL_App_Devices_All)
            .then(function(response) {
                if(!response.data.IsError) {
                    self.items =  JSON.parse(response.data.Data)
                }
            })
            .catch(error => console.log(error));
        },
        remove: function(row) {
            var self = this;
            this.$http.delete(URL_App_Devices_Remove + row.ip)
            .then(function(response) {
                if(!response.data.IsError) {
                    for(var i = 0; i < self.items.length; i++) { 
                        if (self.items[i].ip === row.ip) {
                            self.items.splice(i, 1); 
                        }
                    }
                }
            })
            .catch(error => console.log(error));
        }
    },
    template: `
    <div v-if="items.length > 0" >
        <md-table>
            <md-table-row>
                <md-table-head md-numeric>IP</md-table-head>
                <md-table-head>Port</md-table-head>
            </md-table-row>

            <md-table-row v-for="item in items">
                <md-table-cell md-numeric>{{ item.ip }}</md-table-cell>
                <md-table-cell>{{ item.port }}</md-table-cell>
                <md-table-cell><md-button class="md-raised md-primary" @click="remove(item)">remove</md-button></md-table-cell>
            </md-table-row>
        </md-table>
    </div>
    `,
})