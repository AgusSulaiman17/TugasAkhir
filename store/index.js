export const state = () => ({
  user: null,
  token: null,
  isLoading: false,
});

export const mutations = {
  setUser(state, user) {
    state.user = user;
  },
  setToken(state, token) {
    state.token = token;
  },
  setLoading(state, status) {
    state.isLoading = status;
  }
};

export const actions = {
  created({ commit }) {
    if (process.client) {
      const token = localStorage.getItem('token');
      const user = localStorage.getItem('user');

      if (token) {
        commit('setToken', token);
      }

      if (user) {
        try {
          commit('setUser', JSON.parse(user));
        } catch (error) {
          console.error('Failed to parse user from localStorage:', error);
          localStorage.removeItem('user');
        }
      }
    }
  },

  login({ commit }, { token, user }) {
    commit('setToken', token);
    commit('setUser', user);

    if (process.client) {
      localStorage.setItem('token', token);
      localStorage.setItem('user', JSON.stringify(user));
    }
  },

  logout({ commit }) {
    commit('setToken', null);
    commit('setUser', null);

    if (process.client) {
      localStorage.removeItem('token');
      localStorage.removeItem('user');
    }
  },

  setLoading({ commit }, status) {
    commit('setLoading', status);
  }
};

export const getters = {
  isAuthenticated(state) {
    return !!state.token;
  },
  getUser(state) {
    return state.user;
  },
  isLoading(state) {
    return state.isLoading;
  }
};
