// register the auto-search component
var autosearch = Vue.component('auto-search', {
    data: function () {
        return {
            found: []
        }
    },
    methods: {
        find: function() {
            var self = this;
            self.found = []
            this.$http.get('http://192.168.11.4:8002/v1/autosearch').then(function(response) {
                if(response.status == "200") {
                    self.found = response.data
                }
            })
        }
    },
    template: `
    <div>
        <div>
            <md-button class="md-raised md-primary" @click="find">Find</md-button>
        </div>
        <md-table>
            <md-table-row >
                <md-table-head md-numeric>IP</md-table-head>
                <md-table-head>Port</md-table-head>
            </md-table-row>

            <md-table-row v-for="item in found">
                <md-table-cell md-numeric>{{ item.IP }}</md-table-cell>
                <md-table-cell>{{ item.Port }}</md-table-cell>
            </md-table-row>
        </md-table>
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
        }
    },
    methods: {
        find: function() {
            var self = this;
            self.found = []
            this.$http.get('http://192.168.11.4:8002/v1/manualsearch?ips='+ self.ips + "&ports=" + self.ports).then(function(response) {
                if(response.status == "200") {
                    self.found = response.data
                }
            })
        }
    },
    template: `
    <div>
        <div>
            <md-field md-inline>
                <label>IP range </label>
                <md-input v-model="ips"></md-input>
            </md-field>
        </div>
        <div>
            <md-field md-inline>
                <label>Port range </label>
                <md-input v-model="ports"></md-input>
            </md-field>
        </div>
        <div>
            <md-button class="md-raised md-primary" @click="find">Find</md-button>
        </div>
        <md-table>
            <md-table-row >
                <md-table-head md-numeric>IP</md-table-head>
                <md-table-head>Port</md-table-head>
            </md-table-row>

            <md-table-row v-for="item in found">
                <md-table-cell md-numeric>{{ item.IP }}</md-table-cell>
                <md-table-cell>{{ item.Port }}</md-table-cell>
            </md-table-row>
        </md-table>
    </div>
    `,
})