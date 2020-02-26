const ACQUIRE = true
const RELEASE = false

mutexForF1 = RELEASE
function f1(id) {
  if (mutexForF1 != ACQUIRE) {
    mutexForF1 = ACQUIRE
    setTimeout(() => {
        mutexForF1 = RELEASE
        console.log('WebService called ' + id);
      }, 5000); // image the W.S takes 5 seonds to return. While its waiting no one can call this service again
  } else {
     console.log('f1 - rejected call ' + id);
  }
}


async function asyncCall() {
  console.log('calling');
  var result
  await f1(1);
  await f1(2);
  await f1(3);
  await f1(4);
  await f1(5);
  setTimeout(async () => {
        await f1(6);
      }, 6000);
  // expected output: 'resolved'
}

asyncCall();
