# Devices

Devices are registered with `service.device-registry`.
Various metadata about the devices are stored here, including which controller controls them. This allows the frontend to discover controllers.

Devices have _properties_ and _commands_.

- A property has a value that can be changed, e.g. `power` can be `true` or `false`.
- A commend has no state but performs an action, e.g. `toggleState()`. A command can take arguments.

### Read a device

`GET service.device-controller/device?device_id=foo`

**Response**

```json
{
    "device": {
        ...
    }
}
```

_See `libraries/go/device/device.def` for the device object definition._

### Updating a property

`PATCH service.device-controller/device`

**Request**

- `"device_id":string` the globally unique ID for this device
- `"state":object` a map of property names to new values

The response from reading the device will define the available properties and their types.

**Response**

```json
{
    "device": {
        ...
    }
}
```

_See `libraries/go/device/device.def` for the device object definition._

### Calling a command

`POST service.device-controller/device/cmd`

**Request**

- `"device_id":string` the globally unique ID for this device
- `"command":string` the name of the command to call
- `"args":object` a map of argument names to values

The response from reading the device will define the available commands and their arguments.

**Response**

```json
{}
```

## State providers

Devices do not have dependencies because of the complexity of implementing this in a generic way. This would solve the problem of one device needing another device to be in a particular state before a property can be set, e.g. a light needs the WiFi plug to be on before the brightness can be changed. These kinds of problems can be instead solved by _scenes_.

Devices can, however, have _state providers_. E.g. a WiFi plug might know whether the TV is on or off, and provide that state to the TV's controller which has no way to know on its own.

State providers are only implemented where needed. The state providers controller names are listed as part of the device's metadata. The device's controller, when fetching state, will hit `/provide-state?device_id=<device-id>` on all of the state providers and merge the resulting state together.
