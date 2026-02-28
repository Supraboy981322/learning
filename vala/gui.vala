int main (string[] args) {
    Gtk.init ();

    var window = new Gtk.Window () {
        title = "foo"
    };

    var button = new Gtk.Button.with_label ("DON'T CLICK ME");
    button.clicked.connect (() => {
        button.label = "or do, whatever";
    });

    window.child = button;
    window.present ();

    while (Gtk.Window.get_toplevels ().get_n_items () > 0) {
        GLib.MainContext.@default ().iteration (true);
    }

    return 0;
}
