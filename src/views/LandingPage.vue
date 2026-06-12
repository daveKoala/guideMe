<script setup lang="ts">
// 
</script>

<template>
    <section class="hero">
  <p class="eyebrow">Family travel organiser</p>

  <h1>Stop being the family travel administrator.</h1>

  <p class="hero-text">
    Guide Me keeps your family’s boarding passes, hotel bookings, parking,
    insurance and travel plans in one shared place.
  </p>

  <p class="hero-text">
    Built for parents travelling with twins, triplets and larger families.
  </p>

  <a href="#waitlist" class="button">Join the waitlist</a>
</section>

<section class="problem">
  <h2>Family travel is complicated enough.</h2>

  <p>
    One trip can mean dozens of documents spread across airline apps, emails,
    PDFs, screenshots and booking websites.
  </p>

  <p>
    Guide Me brings everything together so you are not searching for a booking
    reference at the airport.
  </p>
</section>

<section class="features">
  <h2>Everything in one place</h2>

  <div class="feature-grid">
    <article>
      <h3>Store key documents</h3>
      <p>
        Boarding passes, hotel confirmations, parking, insurance, attraction
        tickets and transport details.
      </p>
    </article>

    <article>
      <h3>Organise by traveller</h3>
      <p>
        Link documents to each family member so you can quickly find what belongs
        to who.
      </p>
    </article>

    <article>
      <h3>Build a simple timeline</h3>
      <p>
        See when to leave, when to arrive at the airport, flight times,
        check-in details and return travel.
      </p>
    </article>

    <article>
      <h3>Share with the family</h3>
      <p>
        Invite your partner, grandparents or other travellers so everyone has the
        same trip information.
      </p>
    </article>
  </div>
</section>

<section class="travellers">
  <h2>Designed for real family travel</h2>

  <div class="traveller-list">
    <article>
      <h3>Emma</h3>
      <ul>
        <li>Boarding pass</li>
        <li>Passport reminder</li>
        <li>Insurance details</li>
      </ul>
    </article>

    <article>
      <h3>Oliver</h3>
      <ul>
        <li>Boarding pass</li>
        <li>Passport reminder</li>
        <li>Insurance details</li>
      </ul>
    </article>

    <article>
      <h3>Sophie</h3>
      <ul>
        <li>Boarding pass</li>
        <li>Passport reminder</li>
        <li>Insurance details</li>
      </ul>
    </article>
  </div>
</section>

<section class="story">
  <h2>Built by a parent of twins</h2>

  <p>
    Guide Me started after one family holiday where the important travel
    details were scattered across apps, emails, people, screenshots and PDFs.
  </p>

  <p>
    There had to be a simpler way to keep everyone’s documents, bookings and
    plans together.
  </p>
</section>

<section id="waitlist" class="waitlist">
  <h2>Join the waitlist</h2>

  <p>
    Be first to try Guide Me and help shape a better way to organise family
    travel.
  </p>

  <form>
    <label for="email">Email address</label>
    <input id="email" name="email" type="email" autocomplete="email" required />

    <button type="submit">Join the waitlist</button>
  </form>
  </section>
</template>

<style scoped>
section {
  max-width: var(--layout-max-width);
  margin: 0 auto;
  padding: 64px var(--space-6);
}

section + section {
  border-top: 1px solid var(--color-border);
}

#waitlist {
  display:none
}

h1 {
  font-size: 2.5rem;
  line-height: 1.15;
  font-weight: 700;
  margin-bottom: var(--space-4);
}

h2 {
  font-size: 1.75rem;
  font-weight: 700;
  margin-bottom: var(--space-4);
}

h3 {
  font-size: 1.125rem;
  font-weight: 600;
  margin-bottom: 8px;
}

p {
  margin-bottom: var(--space-4);
  line-height: 1.6;
}

/* Hero */
.hero {
  text-align: center;
  padding-top: 96px;
}

.eyebrow {
  text-transform: uppercase;
  letter-spacing: 0.08em;
  font-size: 0.8rem;
  font-weight: 600;
  color: var(--color-link);
  margin-bottom: var(--space-4);
}

.hero-text {
  font-size: 1.125rem;
  color: #4a5366;
  max-width: 540px;
  margin-inline: auto;
}

/* Buttons */
.button,
button[type="submit"] {
  display: inline-block;
  margin-top: var(--space-4);
  padding: 12px 24px;
  background: var(--color-link);
  color: #fff;
  border-radius: var(--radius-md);
  font-weight: 600;
  cursor: pointer;
  transition: opacity 0.15s ease;
}

.button:hover,
button[type="submit"]:hover {
  opacity: 0.9;
}

/* Grids */
.feature-grid,
.traveller-list {
  display: grid;
  gap: var(--space-6);
  grid-template-columns: repeat(2, 1fr);
}

.feature-grid article,
.traveller-list article {
  padding: var(--space-6);
  background: var(--color-surface);
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
}

.traveller-list {
  grid-template-columns: repeat(3, 1fr);
}

.traveller-list ul {
  list-style: disc;
  padding-left: 20px;
  margin: 0;
}

.traveller-list li {
  line-height: 1.6;
}

/* Waitlist */
.waitlist form {
  display: flex;
  flex-direction: column;
  gap: 8px;
  max-width: 360px;
}

.waitlist label {
  font-weight: 600;
}

.waitlist input {
  padding: 12px;
  border: 1px solid var(--color-border);
  border-radius: var(--radius-md);
}

@media (max-width: 640px) {
  h1 {
    font-size: 2rem;
  }

  .feature-grid,
  .traveller-list {
    grid-template-columns: 1fr;
  }
}
</style>