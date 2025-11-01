On wayland Linux, perhaps we can leverage:

```bash
  interface org.freedesktop.portal.Clipboard {
    methods:
      RequestClipboard(in  o session_handle,
                       in  a{sv} options);
      SetSelection(in  o session_handle,
                   in  a{sv} options);
      SelectionWrite(in  o session_handle,
                     in  u serial,
                     out h fd);
      SelectionWriteDone(in  o session_handle,
                         in  u serial,
                         in  b success);
      SelectionRead(in  o session_handle,
                    in  s mime_type,
                    out h fd);
    signals:
      SelectionOwnerChanged(o session_handle,
                            a{sv} options);
      SelectionTransfer(o session_handle,
                        s mime_type,
                        u serial);
    properties:
      readonly u version = 1;
  };
```

for KDE:
```bash
~ 
 qdbus org.kde.klipper /klipper

signal void org.kde.klipper.klipper.clipboardHistoryUpdated()
method void org.kde.klipper.klipper.clearClipboardContents()
method void org.kde.klipper.klipper.clearClipboardHistory()
method QString org.kde.klipper.klipper.getClipboardContents()
method QString org.kde.klipper.klipper.getClipboardHistoryItem(int i)
method QStringList org.kde.klipper.klipper.getClipboardHistoryMenu()
method void org.kde.klipper.klipper.reloadConfig()
method void org.kde.klipper.klipper.saveClipboardHistory()
method void org.kde.klipper.klipper.setClipboardContents(QString s)
method void org.kde.klipper.klipper.showKlipperManuallyInvokeActionMenu()
method void org.kde.klipper.klipper.showKlipperPopupMenu()
signal void org.freedesktop.DBus.Properties.PropertiesChanged(QString interface_name, QVariantMap changed_properties, QStringList invalidated_properties)
method QDBusVariant org.freedesktop.DBus.Properties.Get(QString interface_name, QString property_name)
method QVariantMap org.freedesktop.DBus.Properties.GetAll(QString interface_name)
method void org.freedesktop.DBus.Properties.Set(QString interface_name, QString property_name, QDBusVariant value)
method QString org.freedesktop.DBus.Introspectable.Introspect()
method QString org.freedesktop.DBus.Peer.GetMachineId()
method void org.freedesktop.DBus.Peer.Ping()
```

This can potentially eliminate the flickering problem on wayland, due to opening and closing of windows.