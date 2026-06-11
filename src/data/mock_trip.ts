export const mockTrip = {
  "trip": {
    "id": "trip_2026_07_01_man_bcn",
    "name": "Family trip to Barcelona",
    "type": "outbound",
    "status": "planned",
    "timezone": "Europe/London",
    "created_at": "2026-06-10T12:00:00+01:00",
    "updated_at": "2026-06-10T12:00:00+01:00"
  },

  "party": {
    "lead_passenger": "Dave",
    "passengers": [
      {
        "id": "passenger_1",
        "name": "Dave",
        "type": "adult"
      },
      {
        "id": "passenger_2",
        "name": "Adult 2",
        "type": "adult"
      },
      {
        "id": "passenger_3",
        "name": "Child 1",
        "type": "child"
      },
      {
        "id": "passenger_4",
        "name": "Child 2",
        "type": "child"
      },
      {
        "id": "passenger_5",
        "name": "Child 3",
        "type": "child"
      }
    ]
  },
  "sharing": {
    "owner_edit_url": "/trips/trip_2026_07_01_man_bcn/edit/secret-edit-token",
    "read_only_url": "/trips/trip_2026_07_01_man_bcn/share/secret-share-token",
    "offline_enabled": true
  }
}