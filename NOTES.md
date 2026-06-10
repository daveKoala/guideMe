The idea is a 'guide for your journey'

A new journey asks for the following (air travel)

Flight number
Starting Postcode
Number of people travelling

The APP creates a time line

Start
How you want to get to airport (car, taxi, public transport)
    This would lead on to secondary questions. E.g. Car -> do you need parking, Taxi -> book here, public transport time tables
    If you are being dropped off by friend then reminder of how much ££££ drop off is

Airport info.. how long before flight you should arrive 
    do you have hold luggaguge
    Any over sized bags, prams
    Do you need assiatnce

Then offer Fast Track, loungages

If we have the filght data - schdeuled and actaul departure AND GATE CLOSES

A place for boarding passes uploaded for now or link to wallet?

Give live arriavl time

Then later we can add onward travel



From a conversation:

MVP scope I’d actually build

I’d build this:

1. Add trip
2. Enter flight number or manual flight details
3. Enter home postcode
4. Pick airport
5. Pick travel mode: driving/taxi/train
6. Generate timeline
7. Save/share timeline
8. Update timeline manually if delay/gate changes

Then add:

9. Flight API enrichment
10. Live traffic
11. Parking/lounges affiliate links
12. Notifications
13. Arrival-side destination timeline
Monetisation options

Realistically:

Option 1: affiliate revenue
parking
lounges
fast track
taxis
insurance
eSIMs

Probably easiest, but noisy.

Option 2: freemium
Free:
manual trips, static timeline

Paid:
live updates, notifications, shared family timeline, multiple trips

Harder to sell, but cleaner.

Option 3: B2B-ish

Sell it to:

small travel agents
airport transfer firms
corporate travel teams
hotel concierge services

That is more work, but probably more credible than £2.99/month consumer subscriptions.

My honest view

This is worth building as a side project if you keep it narrow.

The trap is building:

flight tracker
airport guide
parking marketplace
lounge marketplace
train planner
taxi booking app
travel itinerary manager

That becomes huge and boring.

The sharp version is:

“Enter flight + postcode. I’ll tell you when to leave and what happens next.”

That is a good MVP. Useful. Demoable. Commercially extendable. Not too technically risky.

What got me thinking is a shared journey and all the boarding passes we had

That makes the idea better.

The real pain is not just “when do I leave?” It is:

“There are four of us travelling, everything is scattered across emails/apps/PDFs/screenshots, and I need one shared source of truth for the journey.”

That is a stronger product.

Better framing

Not a flight timeline app.

More like:

Shared travel day pack

Or:

Family journey organiser

It combines:

boarding passes
flight details
airport timeline
parking booking
lounge booking
hotel / destination address
passport reminders
who has what
live updates

The key user is not the solo business traveller. It is the person organising travel for a family/group.

Usually one adult ends up being the human API gateway for the entire trip.

The boarding pass angle is strong

This gives the app a concrete use case:

Upload / save boarding passes
Group them by passenger
Make them available offline
Share the journey with family
Attach reminders and timeline events

Example:

Trip: Manchester → Barcelona

Passengers:
- Dave
- Partner
- Child 1
- Child 2
- Child 3

Documents:
- Boarding pass: Dave
- Boarding pass: Partner
- Boarding pass: Child 1
- Boarding pass: Child 2
- Boarding pass: Child 3
- Parking confirmation
- Hotel booking
- Insurance PDF

Timeline:
07:15 Leave house
08:00 Park car
08:20 Terminal
08:45 Security
09:30 Breakfast
10:15 Gate
10:30 Boarding
11:00 Depart

That is much more useful than yet another flight tracker.

MVP becomes clearer

I would build this first:

1. Create a trip
Trip name
Outbound date
Return date
Departure airport
Arrival airport
2. Add passengers
Name
Optional notes

No accounts per passenger yet. Keep it simple.

3. Upload documents
boarding pass PDFs
screenshots
parking confirmations
hotel booking
insurance docs

Tag each document:

Dave - boarding pass
Child 1 - boarding pass
Parking
Hotel
Insurance
4. Generate simple timeline

Manual first. AI/data later.

Leave home
Arrive airport
Parking
Security
Gate
Boarding
Departure
Arrival
Transfer
Hotel
5. Share trip link

Read-only link:

family can open it
documents grouped clearly
timeline visible
works on phone
offline cache later

That is a genuinely useful MVP.

The killer UX

The home screen should not be a dashboard.

It should be:

Next thing:
08:20 - Go to Terminal 1

Important:
Boarding passes
Parking QR code
Passports

Then sections:

Timeline
People
Documents
Flight
Destination

When you are stood in an airport with kids, you do not want a clever app. You want the next thing and the QR code.

What I would avoid early

Avoid trying to parse every boarding pass perfectly.

Airlines format them differently. PDFs are awkward. Screenshots are awkward. Wallet passes are awkward. OCR becomes a rabbit hole.

For MVP:

User uploads document
User labels it
App stores it
App displays it quickly

Later you can add:

detect passenger name
detect flight number
detect departure time
extract QR/barcode

But do not start there.

Data source strategy

There are now two separate products hiding here:

Product A: timeline intelligence

Flight numbers, delays, airport queues, traffic, gate numbers.

Product B: shared journey pack

Documents, people, timeline, reminders, share link.

Product B is easier, cheaper, and probably more useful early.

So I’d start there.

Then add intelligence to make the pack smarter.

Very simple architecture
PWA
API
SQLite/Postgres
Object storage for documents
Background job for reminders/flight refresh

Core tables:

trips
passengers
documents
timeline_items
share_links

That is enough.

Do not build users/accounts first unless needed.

A trip could be created with a magic edit link and a read-only share link.

Edit link: only organiser has it
Share link: family can view

That keeps onboarding low-friction.

Real product shape

This could become:

“A shared travel folder that turns into a live timeline.”

That is better than “travel itinerary app”, because itinerary apps already exist and are too broad.

Your niche is:

family/group airport journeys
boarding passes
parking/lounges/docs
timeline
offline access
shared visibility
Honest assessment

This is more promising than the original version.

The original was data-heavy and API-dependent.

This version has a useful MVP without needing expensive flight/security/gate data.

Build the shared trip pack first. Add smart timeline calculations after. Then add paid integrations once there is an actual journey object to attach them to.






## Key Design Decisions
1. **CSS-only layouts**: No Tailwind/Bootstrap - pure CSS + variables for simplicity
2. **Router metadata**: Routes specify templates via meta.template for flexibility
3. **Global reset**: Tailwind-like preflight ensures consistent baseline across browsers
4. **Named slots**: BasicLayout uses named slots for future multi-region layouts