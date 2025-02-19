const container = {
  hidden: { opacity: 0 },
  show: {
    opacity: 1,
    transition: {
      delayChildren: 0.5,
      staggerChildren: 0.2,
    },
  },
  exit: {
    opacity: 0,
    transition: {
      delayChildren: 0.5,
      staggerChildren: 0.2,
    },
  },
};

const child = {
  hidden: { opacity: 0 },
  show: { opacity: 1 },
  exit: { opacity: 0 },
};

export { container, child };