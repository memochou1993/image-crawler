new Vue({
  el: '#app',
  delimiters: ['${', '}'],
  vuetify: new Vuetify(),
  data: {
    previewed: false,
    loading: false,
    images: [],
    link: '',
    links: ["https://www.google.com/doodles?hl=zh-TW"],
  },
  computed: {
    endpoint() {
      return `/api?links=${this.links.join(',')}`;
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
      if (!this.link) {
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
        behavior: 'smooth'
      });
    },
    refresh() {
      this.setPreviewed(false);
      this.setLoading(false);
      this.setImages([]);
      this.setLink('');
      this.setLinks([]);
    },
    preview() {
      if (this.link) {
        this.addLink();
      }
      if (!this.links.length) {
        return;
      }
      this.setImages([]);
      this.setLoading(true);
      fetch(this.endpoint)
        .then((response) => {
          return response.json();
        })
        .then(({ data }) => {
          this.setPreviewed(true);
          this.setImages(data);
        })
        .catch((err) => {
          console.error(err);
        })
        .finally(() => {
          this.scrollToBottom();
          this.setLoading(false);
        });
    },
    download() {
      //
    },
  },
});
