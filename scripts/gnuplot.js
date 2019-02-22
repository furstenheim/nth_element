const spawn = require('child_process').spawn
const fs = require('fs')
const _ = require('lodash')
require('@geoblink/lodash-mixins').default(_)
const benchmark = fs.readFileSync('a').toString().split('\n')
const benchmarkSplit = _.mapNonNil(benchmark, function (b) {
  if (!b) {
    return null
  }
  const a = b.split('\t')
  const c = a[0].split('/').map(s => s.trim())
  // console.log(b, c, '--', a[2].split(' '))
  return {
    size: parseInt(c[1].match(/Size_(\d+)/)[1]),
    buckets: c[2].match(/Buckets_(\d+)/)[1],
    type: c[3],
    algorithm: c[4].match(/(.*)-4$/)[1].replace('_', ' '),
    iterations: a[1].trim(),
    time: parseInt(a[2].trim().split(' ')[0])
  }
})
const groupedBenchmarks = _.groupBy(benchmarkSplit, function (b) {
  return b.type + '-' + b.buckets
})

_.forEach(groupedBenchmarks, function (benchmarks) {
  const title = `${benchmarks[0].type} - ${benchmarks[0].buckets}.png`
  const out = fs.createWriteStream(title)
  const data = _.map(_.groupBy(benchmarks, 'algorithm'), function (sameAlgBenchmarks, algorithm) {
    const sorted = _.sortBy(sameAlgBenchmarks, 'size')
    return {
      algorithm,
      csv: _.map(sorted, s => `${s.size},${s.time}`).join('\n') + '\ne\n'
    }
  })
  const gnuplot = spawn('gnuplot',
    [
      '-p',
      '-e',
      `set datafile separator ','; set term png; set logscale xy; plot  ${_.map(data, d => `'-' using 1:2 title "${d.algorithm}" with linespoints`).join(', ')}`
    ])
  gnuplot.stdout.pipe(out)
  gnuplot.stderr.on('data', function (e) {
    console.log(title, e.toString())
  })
  _.forEach(data, function ({csv}) {
    gnuplot.stdin.write(csv)
  })
  gnuplot.stdin.end()
})

/*
const gnuplot = spawn('gnuplot',
  [
    '-p',
    '-e',
    `set datafile separator ','; set term png; plot '-' using 1:2 with linespoints, '-' using 1:3 with linespoints`
  ])

gnuplot.stdout.pipe(out)
// gnuplot.stderr.pipe(out)

/!** gnuplot.stdout.on('data', function (e) {
  console.log(e.toString())
})*!/



const command = `1,2,3
3,5,4
e
1,2,3
3,5,4
e
`

gnuplot.stdin.write(command)
gnuplot.stdin.end()
*/
