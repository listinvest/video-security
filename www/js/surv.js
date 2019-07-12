// register the surv-select component
var survselect = Vue.component('surv-select', {
    props: {
        selected: {
          type: Number,
          required: true,
          default: 1,
        },
    },
    data: function () {
        return {
            items: [
                { id: 1, name: "1 ячейка" },
                { id: 2, name: "2 ячейки" },
                { id: 4, name: "4 ячейки" },
            ]
        }
    },
    watch: {
        selected: function (val) {
            this.$parent.countCells = val;
        },
    },
    template: `
        <div>  
            <md-field>
                <label for="select_serv_id">Раскладка</label>
                <md-select v-model="selected" name="select_serv_id" id="select_serv_id">
                    <md-option v-for="item in items" :value="item.id">{{ item.name }}</md-option>
                </md-select>
            </md-field>
        </div>
    `,
})

// register the surv-grid component
var survgrid = Vue.component('surv-grid', {
    props: {
        selected: {
          type: Number,
          required: true,
          default: 1,
        },
    },
    data: function () {
        return {
            cells: [ 
                {  number: 1 }
            ],
        }
    },
    watch: {
        selected: function (val) {

            if (this.cells.length > this.selected) {
                for (var i = this.cells.length - 1; i >= this.selected; i--) {
                    this.cells.splice(i, 1)
                }
            }

            for (var i = 0; i < this.selected; i++) {
                if (this.cells[i] == undefined) {
                    var newCell = {
                        number: 1,
                    } 

                    this.cells.splice(i, 0, newCell)
                }
            }
        },
    },
    template: `
        <div class="div-surv-grid">
            <surv-cell v-for="cell in cells"/>
        </div>
    `,
})

// register the surv-cell component
var survcell = Vue.component('surv-cell', {
    data: function () {
        return {
            srcVideo: '',
            srcInput: 'rtsp://admin:admin@192.168.11.178:554/cam/realmonitor?channel=1&subtype=1&unicast=true&proto=Onvif',
        }
    },
    methods: {
        start: function() {
          
            // ? -> +
            url = this.srcInput.replace("?", "%2B")
            // & -> $
            url = url.replace("&", "%24")

            this.srcVideo = document.location.protocol + "//" + document.location.host + "/v1/videostream?url=" + url
        }
    },
    template: `
    <md-card>
        <md-ripple>
            <md-card-content>
                <div class="div-cell-video">
                    <video :src="srcVideo" autoplay preload="none" class="cell-video"/>
                </div>
            </md-card-content>

            <md-card-actions>
                <table class="table-cell-input">
                <tr>
                    <td>
                        <md-field md-inline>
                            <label>RTSP url</label>
                            <md-input v-model="srcInput"></md-input>
                        </md-field>
                    </td>
                    <td>
                        <md-button class="md-raised md-primary" @click="start">OK</md-button>
                    </td>
                </tr>
                </table>
            </md-card-actions>
        </md-ripple>
    </md-card>
    `
})

// register the surv component
var surv = Vue.component('surv', {
    components: {
        'surv-select': survselect,
        'surv-grid': survgrid,
        'surv-cell': survcell,
    },
    data: function () {
        return {
            countCells: 1,
        }
    },
    template: `
        <div class="div-surv-container">
            <surv-select v-bind:selected="countCells"/>
            <surv-grid v-bind:selected="countCells"/>
        </div>
    `,
})