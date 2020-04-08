new Vue({
  el: '#app',
  delimiters: ['${', '}'],
  vuetify: new Vuetify(),
  data: {
    message: '',
    snackbar: false,
    previewed: false,
    loading: false,
    images: [],
    link: '',
    links: [],
  },
  computed: {
    query() {
      return `links=${this.links.join(',')}`;
    },
  },
  watch: {
    link(to, from) {
      if (to && !from) {
        this.setPreviewed(false);
      }
      if (!to && !!this.images.length) {
        this.setPreviewed(true);
      }
    },
    links(to, from) {
      if (to.length !== from.length) {
        this.setPreviewed(false);
      }
    },
  },
  methods: {
    setMessage(message) {
      this.message = message;
    },
    setSnackbar(snackbar) {
      this.snackbar = snackbar;
    },
    setPreviewed(previewed) {
      this.previewed = previewed;
    },
    setLoading(loading) {
      this.loading = loading;
    },
    setImages(images) {
      this.images = images;
    },
    setLink(link) {
      this.link = link;
    },
    setLinks(links) {
      this.links = links;
    },
    addLink() {
      if (!this.link.trim()) {
        return;
      }
      if (this.links.includes(this.link)) {
        this.setLink('');
        return;
      }
      this.setLinks([this.link, ...this.links]);
      this.setLink('');
    },
    deleteLink(link) {
      this.setLinks(this.links.filter((item) => item !== link));
    },
    scrollToBottom() {
      window.scrollTo({
        top: document.body.scrollHeight,
        behavior: 'smooth',
      });
    },
    refresh() {
      this.setPreviewed(false);
      this.setLoading(false);
      this.setImages([]);
      this.setLink('');
      this.setLinks([]);
    },
    alert() {
      return (err) => {
        const message = err.name === 'AbortError'
          ? 'Request Timeout'
          : 'Network Error';
        this.setMessage(message);
        this.setSnackbar(true);
      };
    },
    preview() {
      this.addLink();
      this.setImages([]);
      this.setLoading(true);
      this.do({
        action: 'preview',
        timeout: 10,
      })
        .then((response) => response.json())
        .then((data) => {
          this.setImages(data);
          this.setPreviewed(true);
        })
        .catch(this.alert())
        .finally(() => {
          this.scrollToBottom();
          this.setLoading(false);
        });
    },
    download() {
      this.setLoading(true);
      this.do({
        action: 'download',
        timeout: 30,
      })
        .then((response) => response.blob())
        .then((data) => {
          const url = window.URL.createObjectURL(data);
          const link = document.createElement('a');
          link.setAttribute('href', url);
          link.setAttribute('download', 'images.zip');
          document.body.appendChild(link);
          link.click();
          document.body.removeChild(link);
          window.URL.revokeObjectURL(url);
        })
        .catch(this.alert())
        .finally(() => {
          this.setLoading(false);
        });
    },
    do({
      action,
      timeout = 5000,
    }) {
      const controller = new AbortController();
      const { signal } = controller;
      setTimeout(() => controller.abort(), timeout * 1000);
      return fetch(`/api/${action}?${this.query}`, { signal });
    },
    validate(link = "") {
      if (link) {
        try {
          const { protocol } = new URL(link);
          return protocol === "http:" || protocol === "https:";
        } catch {
          return false;
        }
      }
      return this.links.every((link) => this.validate(link));
    },
  },
});
