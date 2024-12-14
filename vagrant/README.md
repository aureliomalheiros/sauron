# Prometheus in Vagrant

## Running 🚀

Clone the repository:

```bash
git clone https://github.com/aureliomalheiros/sauron.git
cd sauron/vagrant
```

UP server:

```bash
vagrant up
```

Access server:

```bash
vagrant ssh
```

### Usage and configuration 🌊

- The Prometheus server will be available at http://192.168.56.150:9090 once the VM is up and running.
- To customize the Prometheus configuration, edit the files in the config/ directory and restart the Prometheus service inside the VM.
- In this case I use default settings
- The local `config/` directory is synced to `/etc/prometheus` inside the VM
