const URL_App_Device_Passwords_All = 'http://192.168.11.4:8002/v1/device/auth/all'
const URL_App_Device_Passwords_Add = 'http://192.168.11.4:8002/v1/device/auth/add/'
const URL_App_Device_Passwords_Remove = 'http://192.168.11.4:8002/v1/device/auth/remove/'

// register the device-passwords component
var appDevicePasswords = Vue.component('app-device-passwords', {
    data: function () {
        return {
            login: "",
            password: "",
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
            this.$http.get(URL_App_Device_Passwords_All)
            .then(function(response) {
                if(!response.data.IsError) {
                    self.items =  JSON.parse(response.data.Data)
                }
                else {
                    alert(response.data.ErrorMessage)
                }
            })
            .catch(error => console.log(error));
        },
        add: function() {
            var self = this;
            this.$http.get(URL_App_Device_Passwords_Add + self.login + "/" + self.password)
            .then(function(response) {
                if(!response.data.IsError) {
                    var item = {
                        login: self.login,
                        password: self.password,
                    }
                    self.items.push(item)

                    self.login = ""
                    self.password = ""
                }
                else {
                    alert(response.data.ErrorMessage)
                }
            })
            .catch(error => console.log(error));
        },
        remove: function(row) {
            var self = this;
            this.$http.delete(URL_App_Device_Passwords_Remove + row.login)
            .then(function(response) {
                if(!response.data.IsError) {
                    for(var i = 0; i < self.items.length; i++){ 
                        if (self.items[i].login === row.login) {
                            self.items.splice(i, 1); 
                        }
                     }
                }
                else {
                    alert(response.data.ErrorMessage)
                }
            })
            .catch(error => console.log(error));
        }
    },
    template: `
    <div>
        <div>
            <md-field md-inline>
                <label>Login</label>
                <md-input v-model="login"/>
            </md-field>
        </div>
        <div>
            <md-field md-inline>
                <label>Password</label>
                <md-input v-model="password"/>
            </md-field>
        </div>
        <div>
            <md-button class="md-raised md-primary" @click="add(item)">Add</md-button>
        </div>
        <div v-if="items.length > 0">
            <md-table>
                <md-table-row>
                    <md-table-head md-numeric>Login</md-table-head>
                    <md-table-head>Password</md-table-head>
                </md-table-row>

                <md-table-row v-for="item in items">
                    <md-table-cell md-numeric>{{ item.login }}</md-table-cell>
                    <md-table-cell>{{ item.password }}</md-table-cell>
                    <md-table-cell><md-button class="md-raised md-primary" @click="remove(item)">remove</md-button></md-table-cell>
                </md-table-row>
            </md-table>
        </div>
    </div>
    `,
})