<template>
  <div class="fluid-container">
    <div class="row">
      <div class="col-lg-10 menu">
        <form class="form-inline">
          <div class="form-group mb-2 mr-sm-2">
            <input size=50 v-model="projectName" placeholder="Project name" class="form-control">
          </div>
          <div class="form-group mb-2 mr-sm-2">
            <select class="form-control" v-model="selectedSheetId">
              <option v-for="file in sheets" v-bind:value="file.id">
                {{ file.name }}
              </option>
            </select>
          </div>
          <div class="form-group mb-2 mr-sm-2">
            <div class="input-group">
              <select class="custom-select" v-model="selectedDocId">
                <option v-for="file in docs" v-bind:value="file.id">
                  {{ file.name }}
                </option>
              </select>

              <div class="input-group-append">
                <button type="button" class="btn btn-secondary btn-sm" v-on:click="translate">Translate</button>
                <button type="button" class="btn btn-secondary btn-sm dropdown-toggle dropdown-toggle-split" data-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                  <span class="sr-only">Toggle Dropdown</span>
                </button>
                <div class="dropdown-menu">
                  <span class="dropdown-item-text"><label for="preview" class="inline">Preview</label> <input id="preview" type="checkbox" v-model="previewDoc" /></span>
                  <a class="dropdown-item" :href="docSrc" target="_blank" v-if="selectedDocId">Open</a>
                  <div role="separator" class="dropdown-divider"></div>
                  <button class="dropdown-item" v-on:click="createNewDocument">Create new</button>
                  <button class="dropdown-item" v-on:click="createNewSpreadsheet" v-if="!selectedSheetId">Create new spreadsheet</button>
                </div>
              </div>
            </div>
          </div>
        </form>
      </div>
      <div class="col-lg-2 text-right logout">
        <button v-on:click="logout" class="btn btn-danger" type="button">Logout</button>
      </div>
    </div>
    <div class="row">
      <iframe :src="src" v-if="selectedSheetId" v-bind:class="{ full: !previewDoc }" />
      <iframe :src="docSrc" v-if="selectedDocId && previewDoc" />
    </div>
  </div>
</template>

<script>
import { mapGetters, mapState } from "vuex";

import { AUTH_LOGOUT } from "@/store/actions/auth";
import apiCall from "@/utils/api";

export default {
  name: 'Home',
  data() {
    return {
      sheets: [],
      docs: [],
      previewDoc: false,
      selectedSheetId: null,
      selectedDocId: null,
      get projectName() {
        return localStorage.getItem('projectName') || "";
      },
      set projectName(value) {
        localStorage.setItem('projectName', value);
      }
    }
  },
  created() {
    this.loadSpreadsheets()
    this.loadDocuments()
  },
  methods: {
    loadSpreadsheets: function() {
      self = this;

      apiCall({url: "/sheets/search", method: "POST", data: {filter: this.projectName}})
              .then(resp => {
                self.sheets = resp.data.files;
                self.selectedSheetId = self.sheets[0].id;
              })
              .catch(err => console.log(err));
    },
    loadDocuments: function() {
      self = this;

      apiCall({url: "/documents/search", method: "POST", data: {filter: this.projectName}})
              .then(resp => {
                self.docs = resp.data.files;
                self.selectedDocId = self.docs[0].id;
              })
              .catch(err => console.log(err));
    },
    logout: function() {
      this.$store.dispatch(AUTH_LOGOUT).then(() => this.$router.push("/login"));
    },
    translate: function() {
      self = this;
      apiCall({url: `/sheets/${this.selectedSheetId}/translate/${this.selectedDocId}`, method: "POST"})
              .then(_ => {
                alert("OK")
              })
              .catch(err => {
                alert("ERROR!")
                console.log(err)
              });
    },
    createNewDocument: function() {
      self = this;
      let index = 1
      if(self.docs){
        index = self.docs.length + 1
      }

      let name = `${self.projectName} ${index}`

      apiCall({url: `/documents`, method: "POST", data: {name: name}})
              .then(resp => {
                console.log(resp)
                self.loadDocuments()
              })
              .catch(err => {
                console.log(err)
              });
    },
    createNewSpreadsheet: function() {
      self = this;

      apiCall({url: `/sheets`, method: "POST", data: {name: self.projectName}})
              .then(resp => {
                console.log(resp)
                self.loadSpreadsheets()
              })
              .catch(err => {
                console.log(err)
              });
    }
  },
  computed: {
    src: function() {
      return `https://docs.google.com/spreadsheets/d/${this.selectedSheetId}`
    },
    docSrc: function() {
      return `https://docs.google.com/document/d/${this.selectedDocId}`
    },
    ...mapGetters(["isAuthenticated"])
  }
}
</script>

<style>
  iframe {
    width: 50%;
    height: 100vh;
  }

  iframe.full {
    width: 100%
  }

  label.inline {
    display: inline
  }

  div.menu {
    padding: 2px
  }

  div.logout {
    padding: 2px
  }
</style>
