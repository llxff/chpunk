import apiCall from "@/utils/api";
import axios from "axios"
import {
    AUTH_REQUEST,
    AUTH_ERROR,
    AUTH_LOGOUT,
    OAUTH_REQUEST
} from "../actions/auth";

const state = {
    token: localStorage.getItem("user-token") || "",
    status: "",
    hasLoadedOnce: false
};

const getters = {
    isAuthenticated: state => !!state.token
};

const actions = {
    [AUTH_REQUEST]: ({ commit, dispatch }) => {
        return new Promise((resolve, reject) => {
            commit(AUTH_REQUEST);
            apiCall({ url: "/auth", method: "POST" })
                .then(resp => {
                  window.location = resp.data;
                })
                .catch(err => {
                    commit(AUTH_ERROR, err);
                    localStorage.removeItem("user-token");
                    reject(err);
                });
        });
    },
    [OAUTH_REQUEST]: ({ commit, dispatch }, params) => {
        return new Promise((resolve, reject) => {
            commit(OAUTH_REQUEST);
            apiCall({ url: "/oauth/callback", method: "POST", data: params })
                .then(resp => {
                    commit(OAUTH_REQUEST, resp.data);
                    localStorage.setItem("user-token", resp.data)
                    axios.defaults.headers.common['Authorization'] = resp.data
                    resolve();
                })
                .catch(err => {
                    commit(AUTH_ERROR, err);
                    localStorage.removeItem("user-token");
                    axios.defaults.headers.common['Authorization'] = ""
                    reject(err);
                });
        });
    },
    [AUTH_LOGOUT]: ({ commit }) => {
        return new Promise(resolve => {
            commit(AUTH_LOGOUT);
            localStorage.removeItem("user-token");
            axios.defaults.headers.common['Authorization'] = ""
            resolve();
        });
    }
};

const mutations = {
    [AUTH_REQUEST]: state => {
        state.status = "loading";
    },
    [OAUTH_REQUEST]: (state, token) => {
        state.status = "success";
        state.token = token;
        state.hasLoadedOnce = true;
    },
    [AUTH_ERROR]: state => {
        state.status = "error";
        state.hasLoadedOnce = true;
    },
    [AUTH_LOGOUT]: state => {
        state.token = "";
    }
};

export default {
    state,
    getters,
    actions,
    mutations
};
