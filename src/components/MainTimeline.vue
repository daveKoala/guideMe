<template>
  <div class="container">
    <div class="wrapper">
      <h1>A day in my 'sleepy' life 😅</h1>
      <ul class="sessions">
        <li>
          <div class="time">09:00 AM</div>
          <p>How is it already 9:00? Just how??? 🤯🤯</p>
        </li>
        <li>
          <div class="time">09:05 AM</div>
          <p>Few more minutes of sleep won't do anyone any harm 🤷..</p>
        </li>
        <li>
          <div class="time">09:30 AM</div>
          <p>Get up 🙄</p>
        </li>
        <li>
          <div class="time">1:00 PM</div>
          <p>How can I feel sleepy again?😵</p>
        </li>
        <li>
          <div class="time">01:30 PM</div>
          <p>Lunch time after which sleep just doesn't want to let go of me. 🤝</p>
        </li>
        <li>
          <div class="time">3:30 PM</div>
          <p>Drink the magical chai, it will ward off sleep they said. 🤷‍</p>
        </li>
        <li>
          <div class="time">4:30 PM</div>
          <p>The only time I don't feel sleepy cause it's work out time. I mean walking time. 😹</p>
        </li>
        <li>
          <div class="time">07:00 PM</div>
          <p>Food my tummy needs, sleep my body needs.👿</p>
        </li>
        <li>
          <div class="time">07:30 PM</div>
          <p>My tummy's happy time 🍝</p>
        </li>
        <li>
          <div class="time">10:00 PM</div>
          <p>uh oh!!! fuel low, get some snacks but wait should I just take a quick nap?🤓</p>
        </li>
        <li>
          <div class="time">2:30 PM</div>
          <p>All hail! The time to sleep has finally arrived.😴😴😴😴😴</p>
        </li>
      </ul>
    </div>
  </div>
</template>

<style scoped>
ul,
li {
  list-style: none;
  padding: 0;
}

.container {
  display: flex;
  justify-content: center;
  align-items: center;
  padding: 3rem 0;
  background: linear-gradient(45deg, #209cff, #68e0cf);
}

.wrapper {
  background: #eaf6ff;
  padding: 2rem;
  border-radius: 15px;
}

h1 {
  font-size: 1.1rem;
  font-family: sans-serif;
}

.sessions {
  margin-top: 2rem;
  border-radius: 12px;
  position: relative;
}

li {
  padding-bottom: 1.5rem;
  border-left: 1px solid #abaaed;
  position: relative;
  padding-left: 20px;
  margin-left: 10px;
}

li:last-child {
  border: 0;
  padding-bottom: 0;
}

li::before {
  content: '';
  width: 15px;
  height: 15px;
  background: #fff;
  border: 1px solid #4e5ed3;
  box-shadow: 3px 3px 0 #bab5f8;
  border-radius: 50%;
  position: absolute;
  left: -10px;
  top: 0;
}

.time {
  color: #2a2839;
  font-family: 'Poppins', sans-serif;
  font-weight: 500;
  font-size: 0.9rem;
  margin-bottom: 0.3rem;
}

p {
  color: #4f4f4f;
  font-family: sans-serif;
  line-height: 1.5;
  margin-top: 0.4rem;
}

@media screen and (max-width: 600px) {
  .time {
    font-size: 0.85rem;
  }

  p {
    font-size: .9rem;
  }
}
</style>