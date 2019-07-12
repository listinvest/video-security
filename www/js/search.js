const URL_Autosearch = 'http://192.168.11.4:8002/v1/autosearch'
const URL_Manualsearch = 'http://192.168.11.4:8002/v1/manualsearch'

// register the auto-search component
var autosearch = Vue.component('auto-search', {
    data: function () {
        return {
            found: [],
            isProgress: false,  
        }
    },
    methods: {
        search: function() {
            var self = this;
            self.isProgress = true
            self.found = []
            this.$http.get(URL_Autosearch)
            .then(function(response) {
                if(!response.data.IsError) {
                    self.found = JSON.parse(response.data.Data)
                }
                else {
                    alert(response.data.ErrorMessage)
                }
                self.isProgress = false
            })
            .catch(error => console.log(error));
        }
    },
    template: `
    <div>
        <div>
            <md-button class="md-raised md-primary" @click="search">
                Search
                <md-progress-spinner v-if="isProgress" class="md-accent" :md-diameter="10" :md-stroke="3" md-mode="indeterminate"/>
             </md-button>
        </div>
        <div>
            <find-devices v-bind:found="found"/>
        </div>
    </div>
    `,
})

// register the manual-search component
var manualsearch = Vue.component('manual-search', {
    data: function () {
        return {
            ips: "192.168.11.180-192.168.11.185, 192.168.11.186",
            ports: "80",
            found: [],     
            isProgress: false,       
        }
    },
    methods: {
        search: function() {
            var self = this;
            self.isProgress = true
            self.found = []
            this.$http.get(URL_Manualsearch + '?ips='+ self.ips + "&ports=" + self.ports)
            .then(function(response) {
                if(!response.data.IsError) {
                    self.found = JSON.parse(response.data.Data)
                }
                else {
                    alert(response.data.ErrorMessage)
                }
                self.isProgress = false
            })
            .catch(error => console.log(error));
        }
    },
    template: `
    <div>
        <div>
            <md-field md-inline>
                <label>IP range</label>
                <md-input v-model="ips"/>
            </md-field>
        </div>
        <div>
            <md-field md-inline>
                <label>Port range</label>
                <md-input v-model="ports"/>
            </md-field>
        </div>
        <div>
            <md-button class="md-raised md-primary" @click="search">
                Search
                <md-progress-spinner v-if="isProgress" class="md-accent" :md-diameter="10" :md-stroke="3" md-mode="indeterminate"/>
            </md-button>
        </div>
        <div>
            <find-devices v-bind:found="found"/>
        </div>
    </div>
    `,
})

// register the devices component
var findDevides = Vue.component('find-devices', {
    props: {
        found: {
          type: Array,
        },
    },
    methods: {
        add: function(row) {
            var self = this;

            this.$http.get(URL_App_Devices_Add + row.ip + '/' + row.port)
            .then(function(response) {
                if(response.data.IsError) {
                    alert(response.Data.ErrorMessage)
                    return
                }

                for(var i = 0; i < self.found.length; i++) { 
                    if (self.found[i].ip === row.ip && self.found[i].port === row.port) {
                        self.found.splice(i, 1); 
                    }
                }
            })
            .catch(error => console.log(error));
        }
    },
    template: `
        <div v-if="found.length > 0">
            <md-table>
                <md-table-row>
                    <md-table-head md-numeric>IP</md-table-head>
                    <md-table-head>Port</md-table-head>
                </md-table-row>

                <md-table-row v-for="item in found">
                    <md-table-cell md-numeric>{{ item.ip }}</md-table-cell>
                    <md-table-cell>{{ item.port }}</md-table-cell>
                    <md-table-cell><md-button class="md-raised md-primary" @click="add(item)">add</md-button></md-table-cell>
                </md-table-row>
            </md-table>
        </div>
    `,
})