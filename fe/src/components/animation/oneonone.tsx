const container = (reverse: boolean = false) => ({
  hidden: { opacity: 0 },
  show: {
    opacity: 1,
    transition: {
      delayChildren: 0.5,
      staggerChildren: 0.2,
      staggerDirection: reverse ? -1 : 1,
    },
  },
  exit: {
    opacity: 0,
    transition: {
      delayChildren: 0.5,
      staggerChildren: 0.2,
      staggerDirection: reverse ? -1 : 1,
    },
  },
});

const child = {
  hidden: { opacity: 0 },
  show: { opacity: 1 },
  exit: { opacity: 0 },
};

export { container, child };