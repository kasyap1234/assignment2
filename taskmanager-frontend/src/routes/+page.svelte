<script>
 let src1="../../static/background.jpg"
</script>

<div class="flex flex-col items-center justify-center min-h-screen bg-gray-50">
  <header class="text-center p-8">
    <h1 class="text-5xl font-extrabold text-gray-900">Elevate Your Productivity</h1>
    <p class="text-xl text-gray-600 mt-4">Discover the ultimate task management tool designed for teams and individuals alike.</p>
  </header>

  <section class="container mx-auto px-4 py-12">
    <div class="flex flex-wrap justify-center items-start mb-20">
      <div class="w-full md:w-1/2 lg:w-1/3 p-4">
        <img src="/hero.jpg" alt="Task Management Hero Image" class="rounded-lg shadow-xl w-full h-64 object-cover" />
      </div>
      <div class="w-full md:w-1/2 lg:w-2/3 p-4">
        <h2 class="text-3xl font-semibold text-gray-900">Organize, Collaborate, and Grow</h2>
        <p class="mt-4 text-gray-600">Our app simplifies task management, enabling you to focus on what truly matters. With intuitive features and seamless collaboration tools, boost your productivity today.</p>
        <button class="mt-8 bg-blue-500 hover:bg-blue-700 text-white font-semibold px-6 py-3 rounded-lg transition duration-300 ease-in-out transform hover:-translate-y-1 hover:scale-110">Get Started</button>
      </div>
    </div>

    <div class="w-full">
      <h2 class="text-2xl font-semibold text-center text-gray-900 mb-8">Features That Empower</h2>
      <div class="flex flex-wrap justify-around items-start">
        <div class="w-full sm:w-1/2 md:w-1/3 lg:w-1/4 xl:w-1/5 p-4">
          <img src="background.jpg" alt="Organize Icon" class="h-36 w-auto mx-auto" />
          <h3 class="text-xl font-semibold mt-4 text-gray-900">Intuitive Organization</h3>
          <p class="mt-2 text-gray-600">Categorize tasks effortlessly with our intuitive drag-and-drop system.</p>
        </div>
        <div class="w-full sm:w-1/2 md:w-1/3 lg:w-1/4 xl:w-1/5 p-4">
          <img src="/Collaboration.jpg" alt="Collaborate Icon" class="h-36 w-auto mx-auto" />
          <h3 class="text-xl font-semibold mt-4 text-gray-900">Seamless Collaboration</h3>
          <p class="mt-2 text-gray-600">Work together seamlessly across teams and projects.</p>
        </div>
        <div class="w-full sm:w-1/2 md:w-1/3 lg:w-1/4 xl:w-1/5 p-4">
          <img src="/progress.jpg" alt="Progress Icon" class="h-36 w-auto mx-auto" />
          <h3 class="text-xl font-semibold mt-4 text-gray-900">Track Progress</h3>
          <p class="mt-2 text-gray-600">Monitor progress with real-time updates and insights.</p>
        </div>
      </div>
    </div>
  </section>

  <footer class="bg-blue-500 text-white py-8 mt-8">
    <div class="container mx-auto px-4 text-center">
      <p>&copy; 2024 Task Management App. All rights reserved.</p>
    </div>
  </footer>
</div>

<style>
  /* Tailwind CSS or custom styles */
 .container {
    max-width: 1200px;
    margin-left: auto;
    margin-right: auto;
  }
</style>