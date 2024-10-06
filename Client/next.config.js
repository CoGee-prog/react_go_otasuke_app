/** @type {import('next').NextConfig} */
const { merge } = require('webpack-merge');

const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
	distDir: "./.next",
	env: {
		REACT_APP_FIREBASE_API_KEY: process.env.REACT_APP_FIREBASE_API_KEY,
		REACT_APP_FIREBASE_AUTH_DOMAIN: process.env.REACT_APP_FIREBASE_AUTH_DOMAIN,
		REACT_APP_FIREBASE_PROJECT_ID: process.env.REACT_APP_FIREBASE_PROJECT_ID,
		REACT_APP_FIREBASE_STORAGE_BUCKET: process.env.REACT_APP_FIREBASE_STORAGE_BUCKET,
		REACT_APP_FIREBASE_MESSAGING_SENDER_ID: process.env.REACT_APP_FIREBASE_MESSAGING_SENDER_ID,
		REACT_APP_FIREBASE_APP_ID: process.env.REACT_APP_FIREBASE_APP_ID,
		REACT_APP_FIREBASE_MEASUREMENT_ID: process.env.REACT_APP_FIREBASE_MEASUREMENT_ID,
	},
	async rewrites() {
    return [
      {
        source: "/__/auth/:path*",
        destination: `https://${process.env.REACT_APP_FIREBASE_PROJECT_ID}.firebaseapp.com/__/auth/:path*`,
      },
    ];
  },
}

module.exports = nextConfig
